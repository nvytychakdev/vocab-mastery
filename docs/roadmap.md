# English Learning Platform — Development Roadmap

This document outlines the phased development plan for the platform. Each phase contains clear goals, core ideas, and implementation notes.

## Phase 1 — Core Product Foundation (MVP)
### 🎯 Goal

Build the fundamental learning experience that immediately delivers value to users:
- Learn vocabulary
- Read short texts
- Track progress & maintain streak
- View everything from a unified dashboard
- This phase must be fast, simple, and extensible.

### 💡 Key Ideas
1. **Vocabulary Learning (Core Feature)**
- Word entries: { word, translation, examples, synonyms, partOfSpeech }
- User dictionaries + collections
- Flashcards + quizzes (4-choice, typing)
- Basic SRS (Again → 1 day, Hard → 2 days, Good → 4 days, Easy → 7 days)
- Word-of-the-day widget

2. **Reading Module**

- Short texts categorized by level (A1–B1 for start)
- Built-in dictionary: click word → popup definition
- Comprehension questions (3–6 per text)
- Reading streak and history

3. **Progress / Stats / Daily Goals**

- Daily target (words, reading, quizzes)
- Streak tracking
- XP-like points or progress bars
- SRS-based review queue
- Dashboard summarizing:
- Learned words
- Due reviews
- Today’s reading
- Streak
- Recommendations

4. Dashboard

- Single entry point for the user
- Cards: “Continue learning”, “Review words”, “Daily progress”, “Recommended reading”
- Motivational, simple, not overwhelming

### 🛠 Implementation Details

**Backend**

- Core tables: `words`, `user_words`, `reading_texts`, `reading_questions`, `progress_stats`, `user_settings`
- SRS scheduler runs on-demand when user requests daily tasks
- CRUD for vocabulary collections and progress tracking

**Frontend (Angular)**

- Vocabulary pages: list → word → training
- Reading view with split panel (text + dictionary)
- Dashboard using component cards
- Reusable quiz engine for vocab & reading

**Data Preparation**

- Initial vocabulary set (500–1000 common words)
- 30–50 short reading texts with simple questions
- Minimal, but high-quality content to bootstrap the platform

## Phase 2 — Expansion: Grammar & Reading Depth
### 🎯 Goal

Build structured grammar learning and expand reading content.
This phase deepens the educational value and turns the platform into a full learning tool.

### 💡 Key Ideas
1. Grammar Lessons

- Small, modular grammar lessons:
- Explanation (short, simple, clean)
- Examples + common mistakes
- Tip boxes (“Avoid this mistake…”)
- CEFR grouping (A1, A2, B1)
- Examples of early topics:
- Present Simple
- Articles (a/an/the)
- Simple Past
- Modals (can, should)

2. Grammar Exercises

- Fill-the-gap
- Choose correct sentence
- Rewrite sentence
- Fix the error (“Find the mistake”)
- Match rule to example
- Lessons must directly link to quizzes.

3. Reading Expansion

- Add 100+ new texts across levels
- Add longer stories (2–5 minutes reading)
- Add categories: travel, work, school, daily life
- Improve question variety (ordering, true/false, sentence completion)

4. Content Tools

- Build small internal tools to simplify content creation:
- Grammar lesson editor
- Reading text + question editor
- Vocabulary import tool
- This accelerates content creation and consistency.

### 🛠 Implementation Details
**Backend**

- New tables: grammar_topics, grammar_lessons, grammar_exercises, exercise_attempts
- Content management endpoints
- Link vocabulary to reading (optional)

**Frontend**

- Grammar page with: list → topic → lesson → exercises
- Exercise engine building on existing quiz logic
- Admin pages (if needed) for importing/editing texts

Data Preparation

- Write grammar topics and lessons
- Prepare examples and exercises manually
- Expand reading dataset
- This phase is content-heavy, less about coding.

## Phase 3 — Future Features: Listening & Speaking
### 🎯 Goal

Define future expansions without implementing them yet.
These are long-term differentiating features but not needed for initial success.

### 💡 Key Ideas
1. Listening Module

- Short audio clips (30–60 sec)
- Transcript toggle (show/hide)
- Comprehension questions
- Listening-to-type (simple dictation)
- Categorized difficulty and themes
- Initial source can be text-to-speech (TTS).

2. Speaking Module

(Concept only — no AI or implementation now)
Ideas for future:
- User records voice → system shows transcript
- Compare user recording with native sample
- Pronunciation practice with phonetics (IPA)
- Speaking prompts (simple ones like: “Describe your breakfast.”)

3. Lesson Integration

In the future:

- Vocabulary should link to audio clips
- Reading texts should have “Listen to this text”
- Grammar should have “Speak this example sentence”
