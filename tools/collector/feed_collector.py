import json

import helpers.db as db
from dotenv import load_dotenv

load_dotenv('../.env')

inputs = [
    {"level": "A1", "dictionary_id": "fbdff10e-2d36-4387-be01-489b917a36e3"},
    {"level": "A2", "dictionary_id": "755806b2-6e0d-41aa-8620-63b701747bc2"},
    {"level": "B1", "dictionary_id": "4366ced1-b886-4b20-999c-4fc498d69e60"},
    {"level": "B2", "dictionary_id": "1420a3d8-b790-4224-93a8-44bbb3c1218e"},
    {"level": "C1", "dictionary_id": "1674674d-4b71-4198-aed2-11b383975c74"},
    {"level": "C2", "dictionary_id": "b106e1e7-910b-4a2d-960b-ec05f27d1978"},
]
data = {}

def fill_words(cur, data, dictionary_id):
    try:
        for entry in data:
            # 1. main word
            word_id = db.upsert_word(cur, entry["word"])
            word_dictionary_id = db.upsert_dictionary_word(cur, dictionary_id,  word_id)

            for meaning in entry["meanings"]:
                print(meaning)
                # 2. part of speech
                pos_id = db.upsert_part_of_speech(cur, meaning["part_of_speech"])

                # 3. meaning
                meaning_id = db.insert_meaning(
                    cur,
                    word_id,
                    meaning["definition"],
                    pos_id
                )

                # 4. example
                db.insert_example(cur, meaning_id, meaning["example"])

            print(f"✅ Word '{entry['word']}' committed")
    except Exception as e:
        print(f"❌ Failed to ingest '{data['word']}': {e}")

def fill_synonyms(cur, data):
    for entry in data:
        for meaning in entry["meanings"]:
            meaning_id = db.get_meaning(cur, meaning["definition"])
            # 5. synonyms
            for syn in meaning["synonyms"]:
                existing_word_id = db.get_word(cur, syn)
                if not existing_word_id:
                    syn_word_id = db.upsert_word(cur, syn)
                else:
                    syn_word_id =existing_word_id 
                db.insert_synonym(cur, meaning_id, syn_word_id)
                print(f"✅ Synonym '{syn}' committed")

# ---------- MAIN INGESTION ----------

def run():
    conn = db.connect()

    with conn:
        with conn.transaction():
            with conn.cursor() as cur:
                for input in inputs:
                    print(input)
                    with open(f"../seeds/{input['level']}-output-v2.json", "r", encoding="utf-8") as f:
                        data[input['level']] = json.load(f)
                        fill_words(cur, data[input["level"]], input['dictionary_id']) 
                for input in inputs:
                    fill_synonyms(cur, data[input['level']]) 
    print("✅ Word ingested successfully")

if __name__ == "__main__":
    run()