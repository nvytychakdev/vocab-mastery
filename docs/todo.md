# Todo

## High priority

* [x] FE: Setup Tailwind and Prettier
* [x] FE: Create auth pages for the app: Login, Register
* [x] FE: Interface to display dictionaries
* [x] FE: Interface to display words. Ability to add words and see the translations
* [x] FE: Add buttons and links UI components
* [x] FE: Add inputs UI componnets
* [x] FE: Add menus UI components
* [x] FE: Add cards UI components
* [x] FE: Add toast UI service
* [x] FE: Add "Email verification" flow to the auth 
* [x] Server: Add "Email verification" flow to the auth
* [x] FE: Add "Email already verified" error handling page 
* [x] Server: Fix local sign-in attempt for OAuth only users 
* [x] Server: Basic configuration (API keys, connections, etc.)
* [x] Server: Setup Postgres docker container and connect to the server
* [x] Server: Setup `User` table schema in DB 
* [x] Server: Setup `Dictionaries` table schema in DB 
* [x] Server: Add pagination support for all GET/POST requests
* [x] Server: Add `include` fields functionality to GET requests
* [x] Server: Add `include` fields functionality to GET lists requests
* [x] Server: Add sort support for all GET/POST requests
* [ ] Server: Add filters support for all GET/POST requests
* [x] Server: Refactor pagination to be reusable for all queries
* [x] Server: Refactor sort to be reusable for all queries
* [x] Server: Setup `Words` table schema in DB 
* [x] API: Basic auth rotues `/sign-in`, `/sign-out`, `/register`, `/refresh`
* [x] API: Advanced OAuth (google sign in support)
* [x] API: CRUD for Dictionaries 
* [x] API: CRUD for Words and Translations
* [x] FE: Implement Dictionaries API integration with UI
* [ ] FE: Loading and page transitions improvements. Move data fetching into the component (avoid resolvers). Dispaly loading states for data fragments loaded from the server.
* [ ] FE: Implement words list smart pagination on scroll. Wrap all words with autosize virtual scroll, implement segments fetching one by one.
* [x] Server: Fill database with all levels of English. A1-2, B1-2, C1-2.
* [ ] API: `/words` filter by `dictionaryId` query params.
* [ ] API: Streamline words creation. Populate meanings, examples and relations between them.
* [ ] API: Add existing words to personal dictionary.
* [ ] Server: Review all meanings, make sure there are no ducplicated, and all of them are precisely described. Collect new meanings/words with better LLM as v2 seeds.
* [ ] Server: Review all synonyms. Make sure during word's creation, all synonyms will be linked back to original word and vice versa. 
* [ ] Server: Populate all words with translations to Russian language.
* [ ] API: Quiz flashcards API. 

## Low Priority:

* [ ] Server: Add coverage for Auth tests on server
* [ ] Server: Add coverage for Dictionaries tests on server
* [x] Server: Refactor Auth to rely on Squirrel
* [ ] Server: Add API docs
* [ ] FE: Implement Words + Translations API integration with UI
* [ ] Server: Translation autocomplete API with Google Translate and Dictonary integration