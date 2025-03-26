<a id="readme-top"></a>



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">History Hunters BE</h3>

  <p align="center">
    An awesome README template to jumpstart your projects!
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#run">Run</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
      <ul>
          <li><a href="#accessing-players">Accessing Players</a></li>
          <li><a href="#accessing-stages">Accessing Stages</a></li>
          <li><a href="#accessing-player-sessions">Accessing Player Sessions</a></li>
          <li><a href="#accessing-questions">Accessing Questions</a></li>
          <li><a href="#accessing-Answers">Accessing Answers</a></li>
      </ul>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
</details>

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- ABOUT THE PROJECT -->
## About the project


<!-- GETTING STARTED -->
## Getting Started
To get started with this project, follow the steps below:

1. **Clone the repository**:

```bash
  git clone https://github.com/your-username/your-repo-name.git
  cd your-repo-name
```

2.	**Install dependencies**:
Make sure you have Go installed, then run the following command to install all necessary dependencies:

```bash
  go mod tidy
```

3. **Set up environment variables**:
Create a .env file in the root of the project and add the necessary environment variables (like your database credentials, etc.):

```sql
  DB_HOST=localhost
  DB_PORT=5432
  DB_NAME=your_db_name
  DB_USER=your_db_user
  DB_PASSWORD=your_db_password
```
4. **Run database migrations**:
To set up the database schema, run the following migrations to create the necessary tables:

```bash
  go run cmd/migrate/main.go
```

### Prerequisites

This is application is run on go version 1.24.0.
- Use [this link](https://go.dev/doc/install) to learn more about installing Go locally.

### Installation

<!-- _Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/github_username/repo_name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```
5. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin github_username/repo_name
   git remote -v # confirm the changes
   ``` -->

### Run

To start a local HTTP server you can run this command in your terminal.

```bash
  go run cmd/api/main.go
```
You will see this message logged in the terminal if this step is successful.

```bash
  Server is running on port 8080
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage
### Accessing Players
#### **Create a Player**
**Endpoint**: `POST /players`
- Creates a new player.
- Request Body
```json
{
  "email": "player@example.com",
  "password_digest": "hashed_password",
  "avatar": "avatar_image_url"
}
```
- Response
 - **Status**: `201 Created`
 - **Body**:
 ```json
 {
    "id": 4,
    "email": "player@example.com",
    "password_digest": "hashed_password",
    "avatar": "avatar_image_url",
    "score": 0,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
}
  ```

#### **Get All Players**
**Endpoint**: `GET /players`
- Retrieves a list of all players.

- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
[
    {
        "id": 1,
        "email": "player1@example.com",
        "password_digest": "hashed_password",
        "avatar": "avatar_image_url",
        "score": 10,
        "created_at": "2025-01-01T00:00:00Z",
        "updated_at": "2025-01-01T00:00:00Z"
    },
    {
        "id": 2,
        "email": "player2@example.com",
        "password_digest": "hashed_password",
        "avatar": "avatar_image_url",
        "score": 20,
        "created_at": "2025-01-01T00:00:00Z",
        "updated_at": "2025-01-01T00:00:00Z"
    }
]
  ```

#### **Get One Player**
**Endpoint**: `GET /players{id}`
- Retrieves a single player's details by id.
- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
{
      "id": 1,
      "email": "player1@example.com",
      "password_digest": "hashed_password",
      "avatar": "avatar_image_url",
      "score": 10,
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
  }
```

#### **Update a Player**
**Endpoint**: `PATCH /player{id}`
- Updates a player’s details by id.
- Request Body
```json
{
  "email": "updated_player@example.com",
  "password_digest": "new_hashed_password",
  "avatar": "new_avatar_image_url",
  "score": 50
}
```
- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
 {
   "id": 1,
   "email": "updated_player@example.com",
   "password_digest": "new_hashed_password",
   "avatar": "new_avatar_image_url",
   "score": 50,
   "created_at": "2025-01-01T00:00:00Z",
   "updated_at": "2025-01-01T00:00:00Z"
}
  ```
#### **Delete a Player**
**Endpoint**: `DELETE /players/{id}`
- Deletes a player by id.
- Response
 - **Status**: `204 No Content`

### Accessing Stages

#### **Create a Stage**
**Endpoint**: `POST /stages`
- Creates a new stage.
- Request Body
```json
{
    "title": "Hard Stage",
    "background_img": "stage1.png",
    "difficulty": 3
}
```
- Response
 - **Status**: `201 Created`
 - **Body**:
 ```json
{
    "id": 9,
    "title": "Extra Hard Stage",
    "background_img": "stage1.png",
    "difficulty": 3,
    "created_at": "2025-03-25T14:03:20.301018-04:00",
    "updated_at": "2025-03-25T14:03:20.301018-04:00"
}
  ```

#### **Get All Stages**
**Endpoint**: `GET /stages`
- Retrieves a list of all stages.

- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
[
    {
        "id": 1,
        "title": "Easy Stage",
        "background_img": "stage1.png",
        "difficulty": 1,
        "created_at": "2025-03-24T16:00:58.00178Z",
        "updated_at": "2025-03-24T16:00:58.00178Z"
    },
    {
        "id": 2,
        "title": "Medium Stage",
        "background_img": "stage2.png",
        "difficulty": 2,
        "created_at": "2025-03-24T16:00:58.00178Z",
        "updated_at": "2025-03-24T16:00:58.00178Z"
    },
    {
        "id": 3,
        "title": "Hard Stage",
        "background_img": "stage2.png",
        "difficulty": 3,
        "created_at": "2025-03-24T16:00:58.00178Z",
        "updated_at": "2025-03-24T16:00:58.00178Z"
    },
    {
        "id": 4,
        "title": "Edited Stage",
        "background_img": "stage1.png",
        "difficulty": 1,
        "created_at": "2025-03-24T16:02:01.842038Z",
        "updated_at": "2025-03-25T14:01:02.950661Z"
    }
]
  ```

#### **Get One stage**
**Endpoint**: `GET /stage{id}`
- Retrieves a single stage's details by id.
- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
{
        "id": 3,
        "title": "Hard Stage",
        "background_img": "stage2.png",
        "difficulty": 3,
        "created_at": "2025-03-24T16:00:58.00178Z",
        "updated_at": "2025-03-24T16:00:58.00178Z"
    }
```

#### **Update a Stage**
**Endpoint**: `PATCH /stage{id}`
- Updates a stage’s details by id.
- Request Body
```json
{
    "title": "Edited Stage",
    "background_img": "stage1.png",
    "difficulty": 1
}
```
- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
{
    "id": 4,
    "title": "Edited Stage",
    "background_img": "stage1.png",
    "difficulty": 1,
    "created_at": "2025-03-24T16:02:01.842038Z",
    "updated_at": "2025-03-25T14:01:02.950661Z"
}
  ```
#### **Delete a Stage**
**Endpoint**: `DELETE /Stage/{id}`
- Deletes a Stage by id.
- Response
 - **Status**: `204 No Content`

### Accessing Player Sessions
#### **Create a Stage**
**Endpoint**: `POST /stages`
- Creates a new stage.
- Request Body
```json
{
    "title": "Hard Stage",
    "background_img": "stage1.png",
    "difficulty": 3
}
```
- Response
 - **Status**: `201 Created`
 - **Body**:
 ```json
{
    "id": 9,
    "title": "Extra Hard Stage",
    "background_img": "stage1.png",
    "difficulty": 3,
    "created_at": "2025-03-25T14:03:20.301018-04:00",
    "updated_at": "2025-03-25T14:03:20.301018-04:00"
}
  ```

#### **Get All of a PLayer's sessions**

#### **Get One Player Session**
**Endpoint**: `GET /players/{id}/player_session/{session_id}`
- Retrieves a single session details of a player by session id.
- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
{
    "id": 7,
    "player_id": 1,
    "stage_id": 2,
    "lives": 2,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "2025-03-25T16:57:39.69573Z"
}
```

#### **Update a Player's Session**
**Endpoint**: `PATCH /players/{id}/player_session/{session_id}`
- Updates a player’s session details by id.
- Request Body
```json
{
  "player_id": 1,
  "stage_id": 2,
  "lives": 3
}
```
- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
{
    "id": 7,
    "player_id": 1,
    "stage_id": 2,
    "lives": 3,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "2025-03-25T17:54:27.705329Z"
}
  ```
#### **Delete a Player's Session**

### Accessing Questions
#### **Create a Player Session's Question**
**Endpoint**: `POST /players/{id}/player_session/{session_id}/questions`
- Creates a new question.
- Request Body
```json
{
    "Player_session_id": 1,
    "question_text": "Is this a test question?"
}
```
- Response
 - **Status**: `201 Created`
 - **Body**:
 ```json
{
    "ID": 7,
    "PlayerSessionID": 1,
    "QuestionText": "Is this a test question?",
    "CreatedAt": "2025-03-25T19:08:05.702213Z",
    "UpdatedAt": "2025-03-25T19:08:05.702213Z"
}
  ```

#### **Get One Player Session Question**
**Endpoint**: `GET /players/{id}/player_session/{session_id}/questions/{id}`
- Retrieves a single player session question.

- Response
 - **Status**: `200 OK`
 - **Body**:
 ```json
{
    "ID": 7,
    "PlayerSessionID": 1,
    "QuestionText": "Is this a test question?",
    "CreatedAt": "2025-03-25T19:08:05.702213Z",
    "UpdatedAt": "2025-03-25T19:08:05.702213Z"
}
  ```
### Accessing Answers

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Top contributors:

<a href="https://github.com/othneildrew/Best-README-Template/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=othneildrew/Best-README-Template" alt="contrib.rocks image" />
</a>

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Your Name - [@your_twitter](https://twitter.com/your_username) - email@example.com

Project Link: [https://github.com/your_username/repo_name](https://github.com/your_username/repo_name)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
* [Malven's Grid Cheatsheet](https://grid.malven.co/)
* [Img Shields](https://shields.io)
* [GitHub Pages](https://pages.github.com)
* [Font Awesome](https://fontawesome.com)
* [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 