# HOROLOGER — Time-Lock API Server

This is a time-lock API server written in GoLang. It allows clients to send data and specify a unix timestamp of releasing time. The server then returns a retrieval key. Users can later use the key to retrieve the data if the specified time has passed.

## Features

The server provides the following functionalities:

- Create new records with data and release time
- Retrieve records using the retrieval key if the release time has passed
- Delete records using the retrieval key (regardless of whether the specified release time has passed)

## Installation

1. Make sure you have GoLang installed
2. Clone the repository
3. Navigate to the project directory
4. Load the MySQL database from the `horologer.sql`
5. Enter the relevant data into the `config.yml`
6. Run the server using the command `go run .`

## Usage

### Creating a new record

To create a new record, make a `POST` request to the `/strongbox` endpoint with the following parameters:

 - `text`: The data you want to store, as a string.
 - `available_after`: The Unix timestamp after which the data should be released, as an integer.

Here's an example request:

```http
POST /strongbox HTTP/1.1
Content-Type: application/x-www-form-urlencoded

text=example+data&available_after=1646582400
```

And an example JSON response for `200 OK` status code:

```json
{
    "retrieve_key": "c8fda64b22d56cff4e72sD4iNK0KVX-Vnuk9NtuyUZnW9lyoaXAHbcA8HjS8P60="
}
```

### Retrieving and deleting records

To retrieve a record, make a `GET` request to the `/strongbox` endpoint, where `key` query parameter is the key returned by the server when the record was created.

To delete a record, simply make a `DELETE` request to the `/strongbox` endpoint, where `key` query parameter is the retrieve key (just like when getting a record).

An example URL that can be used to retrieve or delete a record:

```
https://server-hostname.com/strongbox?key=c8fda64b22d56cff4e72sD4iNK0KVX-Vnuk9NtuyUZnW9lyoaXAHbcA8HjS8P60=
```

## License

This project is licensed under the MIT License. See the [LICENSE](https://raw.githubusercontent.com/wachttijd/horologer/main/LICENSE) file for details.
