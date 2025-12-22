# Core database contept

```
User
 └── Dictionary (default or personal)
      └── DictionaryWord
           └── Word
                ├── Meanings
                │    ├── Examples
                │    ├── Translations
                │    └── Synonyms
                └── Tags
```

General rules:
- Each user will be able to see list of default dictionaries, split by categories or difficuly
- Each user will be able to create personalized dictionary
- Dictionaries will comtain list of words
- Each word may contain list of meanings (at least one) with coresponding Examples, Translations and Synonyms references
- Each word may contain tags that can be used to categorize words despite of meaning

Future plans:
- Each user may be able to learn words, so the User -> Word relation should be stored, tracking progress of the learning process.
- Each user may write simple notes for some words for better understanding. This should improve learning process, by allowing user to review his own notes as a hint for better understanding what the word means. This feature should not interfere with initial meaning of the word, so the learning weight after revealing the "hint" supposed to be reduced until user is able to proceed without having to peak into the note.
- Introduce words collocations. These are commonly used words next to specified word such as "heavy rain" for the word "rain". 
- Introduce labels for words:
  - slang
  - formal
  - informal
  - offensive
  - academic
  - archaic
- Introduce word forms (pronunciation).

## Dictionary: 

```
dictionaries
--------------------------------------
id                  UUID PK
owner_id            UUID FK -> users.id or NULL
title               text 
level               text 
is_default          boolean
created_at          timestamptz 
```


```
dictonary_words
--------------------------------------
dictionary_id       UUID FK -> dictionaries.id
word_id             UUID FK -> words.id
added_at            timestamptz 
```

## Word: 

```
parts_of_speech
--------------------------------------
id                  UUID PK
code                text
```

```
words
--------------------------------------
id                  UUID PK
word text           NOT NULL
created_at          timestamptz 
```

```
word_meanings
--------------------------------------
id                  UUID PK
word_id             UUID FK -> words.id
definition          text NOT NULL
part_of_speech_id   UUID FK -> parts_of_speech.id 
```


```
word_examples
--------------------------------------
id                  UUID PK
meaning_id          UUID FK -> words.id
text                text NOT NULL
```

```
word_synonyms
--------------------------------------
meaning_id          UUID FK -> word_meanings.id 
synonym_word_id     UUID FK -> words.id
```


## Translations

```
languages
--------------------------------------
code                text (es, ru)
name                text 
```

```

word_translations
--------------------------------------
id                  UUID PK 
meaning_id          UUID FK -> word_meanings.id
language            text FK -> languages.code
translation         text NOT NULL 
```


## Tags

```
tags
--------------------------------------
id                  UUID PK
name                text UNIQUE
```

```

word_tags
--------------------------------------
word_id             UUID FK -> words.id
tag_id              UUID FK -> tags.id
```

# Future database improvements

## Learning 

```

user_words
--------------------------------------
user_id             UUID FK -> users.id
word_id             UUID FK -> words.id
status              text (new, ongoing, mastered)
difficuly           int
last_seen           timestamptz
```

## Usage Labels

```
usage_labels
--------------------------------------
id                  UUID FK -> users.id
label               text NOT NULL
```

```
meaning_labels
--------------------------------------
meaning_id          UUID FK -> word_meanings.id
label_id            UUID FK -> usage_labels.id
```