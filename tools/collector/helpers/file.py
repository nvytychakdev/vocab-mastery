from io import TextIOWrapper
import json
from typing import Callable
import typer
from helpers.error import TranslationError


def write_list_json(
    output_file: str,
    output_error_file: str,
    entries: list[dict[str, str]],
    writer: Callable[[TextIOWrapper, dict[str, str]], bool],
):
    error_records = []

    if not entries or len(entries) == 0:
        return error_records

    try:
        with open(output_file, "w", encoding="utf-8") as f:
            f.write("[\n")
            with typer.progressbar(
                length=len(entries), label="Processing words..."
            ) as progress:
                for entry in entries:
                    try:
                        progress.label = f"Processing: {entry['word']}, errors: ({len(error_records)}/{len(entries)})"
                        processed = writer(f, entry)
                        if not processed:
                            continue
                        f.write(",\n")
                        f.flush()
                    except TranslationError:
                        record = {
                            "word": entry["word"],
                            "meaning_id": entry["meaning_id"],
                            "meaning": entry["meaning"],
                        }
                        error_records.append(record)

                        print(f"\nProblematic record: {record}")

                    finally:
                        progress.update(1)
    finally:
        with open(output_file, "a", encoding="utf-8") as f:
            f.write("]\n")
        with open(output_error_file, "w", encoding="utf-8") as f:
            f.write(json.dumps(error_records))
    return error_records
