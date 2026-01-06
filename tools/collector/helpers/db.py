import os
import uuid
import psycopg
from psycopg.rows import dict_row


def connect():
    return psycopg.connect(
        f"postgresql://{os.getenv('POSTGRES_USER')}:{os.getenv('POSTGRES_PASSWORD')}@localhost:5432/{os.getenv('POSTGRES_DB')}",
        row_factory=dict_row,  # type: ignore
    )


def upsert_dictionary_word(cur, dictionary_id, word_id):
    cur.execute(
        """
      INSERT INTO dictionary_words (dictionary_id, word_id)
      VALUES (%s, %s)
      RETURNING dictionary_id
  """,
        (dictionary_id, word_id),
    )
    return cur.fetchone()["dictionary_id"]


def upsert_word(cur, word):
    cur.execute(
        """
        INSERT INTO words (id, word)
        VALUES (%s, %s)
        RETURNING id
    """,
        (uuid.uuid4(), word),
    )
    return cur.fetchone()["id"]


def get_meaning(cur, definition):
    # 1️⃣ Try to find existing word in this dictionary
    cur.execute(
        """
        SELECT w.id AS id
        FROM word_meanings w
        WHERE w.definition = %s
        LIMIT 1
    """,
        (definition,),
    )

    row = cur.fetchone()
    if row:
        return row["id"]
    return None


def get_word(cur, word):
    # 1️⃣ Try to find existing word in this dictionary
    cur.execute(
        """
        SELECT w.id AS id
        FROM words w
        WHERE w.word = %s
        LIMIT 1
    """,
        (word,),
    )

    row = cur.fetchone()
    if row:
        return row["id"]
    return None


def upsert_part_of_speech(cur, code):
    cur.execute(
        """
        INSERT INTO parts_of_speech (id, code)
        VALUES (%s, %s)
        ON CONFLICT (code) DO UPDATE SET code = EXCLUDED.code
        RETURNING id
    """,
        (uuid.uuid4(), code),
    )
    return cur.fetchone()["id"]


def insert_meaning(cur, word_id, definition, pos_id):
    meaning_id = uuid.uuid4()
    cur.execute(
        """
        INSERT INTO word_meanings (id, word_id, definition, part_of_speech_id)
        VALUES (%s, %s, %s, %s)
    """,
        (meaning_id, word_id, definition, pos_id),
    )
    return meaning_id


def insert_example(cur, meaning_id, text):
    cur.execute(
        """
        INSERT INTO word_examples (id, meaning_id, text)
        VALUES (%s, %s, %s)
    """,
        (uuid.uuid4(), meaning_id, text),
    )


def insert_synonym(cur, meaning_id, synonym_word_id):
    cur.execute(
        """
        INSERT INTO word_synonyms (meaning_id, synonym_word_id)
        VALUES (%s, %s)
        ON CONFLICT DO NOTHING
    """,
        (meaning_id, synonym_word_id),
    )
