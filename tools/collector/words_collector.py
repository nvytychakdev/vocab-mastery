import json
import time
from ollama import ChatResponse, chat
import helpers.prompt as prompt_manager

version = "v3"
level = "A1"
input_file = f"../seeds/{level}-words.txt"
output_file = f"../seeds/{level}-output-{version}.json"
generative_model = "krith/qwen2.5-14b-instruct:IQ3_S"


def get_llm_response(model, prompt):
    response: ChatResponse = chat(
        model=model, messages=[{"role": "user", "content": prompt}]
    )
    return response


def run():
    with open(input_file, "r", encoding="utf-8") as f:
        lines = [line.strip() for line in f if line.strip()]

    print("Starting to request LLM...")
    start = time.time()

    words_with_errors = []
    try:
        with open(output_file, "w", encoding="utf-8") as f:
            f.write("[\n")
            for index, word in enumerate(lines):
                prompt = prompt_manager.get_full_prompt(word, level)
                response = get_llm_response(generative_model, prompt)

                try:
                    if not response.message.content:
                        continue

                    data = json.loads(response.message.content)
                    # data["word"] = word
                    data["level"] = level
                    print(
                        f"{index}/{len(lines)}: ",
                        word,
                        " - Success. Data parsed and added to the file",
                    )
                    f.write(json.dumps(data, ensure_ascii=False) + ",\n")
                    f.flush()
                except json.JSONDecodeError as e:
                    print(
                        f"{index}/{len(lines)}: ",
                        word,
                        "- Error. Failed to parse JSON: ",
                        e,
                    )
                    print("Response from LLM: ", response.message.content)
                    words_with_errors.append(word)

        end = time.time()
        print(f"Successfully finished in {end - start}")
        print("Words with errors:", words_with_errors)
    finally:
        with open(output_file, "a", encoding="utf-8") as f:
            f.write("]\n")


if __name__ == "__main__":
    run()

