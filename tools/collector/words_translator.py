import json
import re
import time
from ollama import ChatResponse, chat
import helpers.prompt as prompt_manager
from helpers import db
from helpers import file
from helpers.error import TranslationError


version = "v3"
language = "Russian"
output_file = f"../seeds/translations-output-{version}.json"
output_error_file = f"../seeds/translations-error-{version}.json"
generative_models = [
    "aya:8b",
    "krith/qwen2.5-14b-instruct:IQ3_S",
    "qwen2.5:7b-instruct-q8_0",
    "llama3.1:8b-instruct-q8_0",
]

CYRILLIC_RE = re.compile(r"^[\u0400-\u04FF\s-]+$")
MAX_RETRIES = 3


def get_llm_response(model, prompt):
    response: ChatResponse = chat(
        model=model,
        messages=[{"role": "user", "content": prompt}],
        options={
            "temperature": 0,
            "top_k": 1,
            "top_p": 1,
        },
    )
    return response


def translate_writer(f, entry):
    prompt = prompt_manager.get_translate_prompt(
        entry["word"], entry["meaning"], language
    )

    for attempt in range(1, MAX_RETRIES + 1):
        response = get_llm_response(generative_models[attempt - 1], prompt)

        if not response.message.content:
            return False

        try:
            words = json.loads(response.message.content)

            if not words or not all(CYRILLIC_RE.match(w) for w in words):
                raise ValueError("Translation is not proper", response.message.content)

            data = {
                "word": entry["word"],
                "meaning": entry["meaning"],
                "meaning_id": entry["meaning_id"],
                "translations": words,
            }
            f.write(json.dumps(data, ensure_ascii=False))
            f.flush()
            return True
        except (json.JSONDecodeError, ValueError, TypeError):
            if attempt < MAX_RETRIES:
                time.sleep(0.5)
                continue
            raise TranslationError(f"Not possible to translate {entry['word']}")

    return False


def get_words(conn):
    words = []
    with conn:
        with conn.transaction():
            with conn.cursor() as cur:
                print("Requested list of all words, it may take some time...")
                words = db.get_words_list(cur)

    if not words or not len(words):
        print("No words found. Exit...")
        return None

    return words


def run():
    print("Connecting to DB...")
    conn = db.connect()
    print("Request list of words...")
    entries = get_words(conn)
    if entries:
        print("Start connection to LLM...")
        file.write_list_json(output_file, output_error_file, entries, translate_writer)
    return
