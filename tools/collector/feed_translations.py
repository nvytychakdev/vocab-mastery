import json

from helpers import db


version = "v3"
language_code = "ru"
input_file = f"../seeds/translations-output-{version}.json"


def get_translations():
    with open(input_file, "r", encoding="utf-8") as f:
        return json.load(f)
    return None


def run():
    conn = db.connect()
    entries = get_translations()

    with conn:
        with conn.transaction():
            with conn.cursor() as cur:
                for entry in entries:
                    for translation in entry["translations"]:
                        db.insert_translation(
                            cur,
                            entry["meaning_id"],
                            language_code,
                            translation,
                        )
    return
