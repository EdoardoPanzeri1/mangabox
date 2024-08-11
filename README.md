# Mangabox

Mangabox is a user-friendly application for manga enthusiasts to search for, track, and manage their manga collections using data fetched from the Jikan API.

![Alt text]([URL-to-your-GIF](https://github.com/EdoardoPanzeri1/mangabox/blob/main/chrome_4nr1BPWlXv.gif?raw=true))

## 📚 Motivation

I have always been a huge fan of the Letterboxd application and wanted something similar for manga enthusiasts. After searching for a comparable platform and finding none that suited my needs, I decided to create Mangabox to fill that gap.

## Features

- **Search and View Manga:** Users can search and view details about their favorite manga.
- **Catalog Management:** Add manga to your personal catalog and update its reading status from "bought" to "read," or delete unwanted entries.
- **User Authentication:** Secure account creation with username, email, and password, and the ability to update your profile.
- **Database Security:** All data is handled securely in a local Dockerized environment.

## 🐳 Installation

1. Ensure you have Docker and Docker-Compose installed.
2. Clone the repository and navigate to the root directory.
3. **Configure Environment Variables:**
   - Create a `.env` file using the following template:

     ```plaintext
     PORT=8080
     REACT_APP_API_URL=http://localhost:8080
     DATABASE_URL=postgres://<POSTGRES_USER>:<POSTGRES_PASSWORD>@db:5432/<POSTGRES_DB>?sslmode=disable
     JIKAN_BASE_URL=https://api.jikan.moe/v3
     POSTGRES_USER=<POSTGRES_USER>
     POSTGRES_PASSWORD=<POSTGRES_PASSWORD>
     POSTGRES_DB=<POSTGRES_DB>
     ```

   - Replace placeholders with your actual credentials and any necessary configuration.
4. Run `docker-compose up --build` to set up and start the project.

## Usage

- After installation, access the Mangabox frontend via your browser by navigating to `http://localhost:3000/`.
- Use the interface to search for manga and manage your catalog.

## 👏 Contributing

- Contributions are welcome! Consider implementing additional features like a wishlist to keep track of manga you plan to read and a social-like network feature to connect with other manga enthusiasts.
- Please follow the contribution guidelines provided in the repository.
