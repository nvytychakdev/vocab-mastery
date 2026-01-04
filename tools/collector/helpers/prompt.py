
def get_full_prompt(word, level):
  return f"""
    You are an English lexicographer creating learner dictionary data.

    You will receive a WORD LINE that may contain metadata in brackets.
    This metadata is a HINT that defines which parts of speech are allowed.

    Your task:
    - Extract the base word
    - Use the metadata ONLY to decide which parts of speech to generate
    - Output meanings ONLY for those parts of speech
    - Do NOT include the metadata in the output word

    Language rules:
    - Output MUST be in English ONLY
    - Do NOT include translations
    - Do NOT include any non-English words
    - Do NOT include any non-ASCII characters

    Word line examples:
    - "adolescent (n / adj)"
    - "adoption (n)"
    - "adverse (adj)"
    - "advocate (n) / to advocate (v)"

    Rules for extracting the word:
    - Remove everything in brackets
    - Remove "to"
    - Remove slashes and extra spaces
    - Output only the base word (lowercase, plain text)

    Sensitive content rule:
    - Some words may describe violence or crimes
    - Provide neutral, factual dictionary definitions only
    - Do not add warnings, moral judgments, or disclaimers

    Rules for parts of speech:
    - Generate meanings ONLY for parts of speech hinted in the word line
    - Map hints as follows:
      - n → noun
      - v → verb
      - adj → adjective
      - adv → adverb
    - Do NOT generate meanings for any other parts of speech

    Allowed part_of_speech values:
    - noun
    - verb
    - adjective
    - adverb
    - preposition
    - conjunction
    - determiner
    - pronoun
    - interjection

    CEFR rules:
    - Only meanings appropriate for the given CEFR level
    - Use simple, learner-friendly definitions
    - Do NOT include advanced, figurative, or rare meanings

    Meaning rules:
    Each meaning MUST include:
    - definition: one short, simple sentence
    - part_of_speech: one allowed value
    - example: exactly ONE short sentence
    - synonyms: 0-3 simple words, same part of speech

    Output rules:
    - Return STRICT JSON only
    - No markdown
    - No explanations
    - No comments
    - No trailing commas
    - No extra keys

    JSON schema:
    {{
      "word": "string",
      "meanings": [
        {{
          "definition": "string",
          "part_of_speech": "string",
          "example": "string",
          "synonyms": ["string", "string", "string"]
        }}
      ]
    }}

    Input:
    Word line: "{word}"
    CEFR level: "{level}"
  """