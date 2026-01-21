# Core database contept

```
User
 ├── Dictionary (default or personal)
 │    └── DictionaryWord
 │         └── Word
 │              ├── Meanings
 │              │    ├── Examples
 │              │    ├── Translations
 │              │    └── Synonyms
 │              └── Tags
 │
 └── Flashcards
      ├── Engagement (1 per user)
      │    └── Reminder / Activity State
      │
      ├── Days (1 per active day)
          └── Sessions (1..N per day)
               └── Attempts (1..N per session)
                    └── Meaning (FK → WordMeaning)
```

General rules:
- Each user will be able to see list of default dictionaries, split by categories or difficulty
- Each user will be able to create personalized dictionary
- Dictionaries will contain list of words
- Each word may contain list of meanings (at least one) with corresponding Examples, Translations and Synonyms references
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
dictionary_words
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
word                text NOT NULL
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
meaning_id          UUID FK -> word_meanings.id
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
user_words_progress
--------------------------------------
id                            UUID PK
user_id                       UUID FK -> users.id
meaning_id                    UUID FK -> word_meanings.id
status                        text -- 'new' | 'learning' | 'review' | 'mastered'
difficulty                    int
times_seen_recall             int
times_correct_recall          int
times_incorrect_recall        int
next_review_at_recall         timestamptz
times_seen_recognition        int
times_correct_recognition     int
times_incorrect_recognition   int
next_review_at_recognition    timestamptz
last_seen_at                  timestamptz
created_at                    timestamptz 
```

### `flashcard_engagement`
#### Purpose

Tracks a user’s long-term engagement state with flashcards.
This table is 1:1 with users and exists only for users who have used flashcards at least once.

It is used for:

- Reminder logic (daily / weekly / disabled)
- Detecting inactivity
- Tracking missed days
- Determining when to send notifications

#### Row lifecycle

Created when:

- User starts their first ever flashcard session

Updated when:

- User starts a flashcard session
- User submits a flashcard answer
- Reminder logic advances (daily → weekly → disabled)

```
flashcard_engagement
--------------------------------------
user_id                       UUID PK FK -> users.id
last_active_at                timestamptz NOT NULL
last_session_date             date
reminder_stage                text NOT NULL -- 'daily' | 'weekly' | 'disabled'
missed_days_count             int NOT NULL DEFAULT 0
next_reminder_at              timestamptz
created_at                    timestamptz DEFAULT now()
updated_at                    timestamptz DEFAULT now()
```

### `flashcard_days`
#### Purpose

Represents a single calendar day of flashcard activity for a user.

Used to:

- Track whether the user studied on a specific day
- Store daily progress and completion
- Support streaks and daily goals
- Row lifecycle

Created when:

- User starts their first flashcard session for a given day

Updated when:

- A session starts or completes
- Flashcard attempts are submitted
- Daily session is completed

```
flashcard_days
--------------------------------------
id                            UUID PK
user_id                       UUID NOT NULL FK -> users.id
date                          date NOT NULL
timezone                      text NOT NULL
started_at                    timestamptz
completed_at                  timestamptz
sessions_count                integer NOT NULL DEFAULT 0
cards_answered                integer NOT NULL DEFAULT 0
cards_correct                 integer NOT NULL DEFAULT 0
created_at                    timestamptz NOT NULL DEFAULT now()
updated_at                    timestamptz NOT NULL DEFAULT now()

UNIQUE (user_id, date)
```

### `flashcard_sessions`
#### Purpose

Represents a single learning session within a day.

Used to:

- Allow multiple sessions per day
- Resume unfinished sessions
- Track progress independently of daily limits

#### Row lifecycle

Created when:

- User starts a new flashcard session

Updated when:

- Cards are answered
- Session is completed or abandoned

```
flashcard_sessions
--------------------------------------
id                                 UUID PK
user_id                            UUID FK -> users.id
flashcard_day_id                   UUID FK -> flashcard_days.id
current_meaning_id                 UUID FK -> word_meanings.id
current_meaning_translation_id     UUID FC -> word_translations.id
started_at                         timestamptz
ended_at                           timestamptz
cards_total                        int
cards_completed                    int
```

### `flashcard_attempts`
#### Purpose

Stores individual answers to flashcards.

Used for:

- Learning progress
- Accuracy calculations
- Spaced repetition / difficulty tuning
- Analytics
- Row lifecycle

Created when:

- User submits an answer for a flashcard

**Never updated (append-only).**

```
flashcard_attempts
--------------------------------------
id                            UUID PK
session_id                    UUID FK -> flashcard_sessions.id
meaning_id                    UUID FK -> word_meanings.id
direction                     text NOT NULL  -- 'recall' or 'recognition'
prompt_language               text NOT NULL  -- 'EN' or user's native language, the language of the question
answer_language               text NOT NULL  -- 'EN' or user's native language, the language of the answer options
is_correct                    boolean
response_time_ms              int
created_at                    timestamptz
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