# 🧠 Word Learning & Progress Tracking

This document outlines how user learning progress is tracked for each word in the app using a spaced repetition-inspired system.

---

## 📦 `word_progress` Table

Tracks how well a specific user remembers a word based on quiz performance.

### Schema

| Column              | Type      | Description                                 |
|---------------------|-----------|---------------------------------------------|
| `id`                | UUID      | Internal ID                                 |
| `user_id`           | UUID      | Foreign key to `users`                      |
| `word_id`           | UUID      | Foreign key to `words`                      |
| `correct_attempts`  | INT       | Total correct answers by the user           |
| `incorrect_attempts`| INT       | Total incorrect answers by the user         |
| `streak`            | INT       | Number of correct answers in a row          |
| `last_seen_at`      | TIMESTAMP | Last time this word was shown in quiz       |
| `next_due_at`       | TIMESTAMP | When the word should be shown again         |
| `easiness_factor`   | FLOAT     | Memory strength rating (SM-2 style)         |
| `is_learned`        | BOOLEAN   | Whether the word is considered "learned"    |
| `created_at`        | TIMESTAMP | When tracking started                       |

---

## 🧮 Learning Logic

We use a simplified [SM-2 algorithm](https://www.supermemo.com/en/archives1990-2015/english/ol/sm2) (used by Anki) to model long-term retention.

### When the user answers **correctly**:

- Increment `streak`
- Update `easiness_factor`:

> EF = EF + 0.1 - (5 - quality) * (0.08 + (5 - quality) * 0.02)
> EF = max(EF, 1.3)

- Calculate `next_due_at` based on `streak` and `easiness_factor` (e.g. 1d, 3d, 7d, etc.)

### When the user answers **incorrectly**:

- Reset `streak` to 0
- Decrease `easiness_factor`:
> EF = max(EF - 0.2, 1.3)
- Set `next_due_at = NOW() + 1 day`

---

## ✅ Marking a Word as Learned

A word is considered **learned** when:

>streak >= 3 AND easiness_factor >= 2.5

This is stored in `is_learned`, and helps exclude it from future quizzes.

---

## 🎮 Quiz Selection Logic

When generating a quiz for a user:

- Prefer words:
  - Not yet marked `is_learned`
  - Or where `next_due_at <= NOW()`
- Prioritize words with:
  - Recent incorrect answers
  - Low `streak` or low `easiness_factor`

Learned words may still occasionally be shown to support long-term retention.

---

## 🧩 Notes

- `quality` is a value between 0–5 representing how well the user recalled the word:
  - 5 = perfect
  - 3 = hesitant
  - 0 = incorrect
- `easiness_factor` starts at **2.5** and adjusts dynamically.
- All updates to this table occur **after each quiz answer**.