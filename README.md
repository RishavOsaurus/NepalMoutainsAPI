# Nepal Peaks API

This is a Go-based API that provides information about peaks in Nepal. The API allows users to retrieve details about different peaks, search for specific peaks by name, and filter peaks using their unique ID. The data served by the API has been scraped from the official [Nepal Himal Peak Profile](https://nepalhimalpeakprofile.org/peak-profile) website.

## Endpoints

### 1. Get All Peaks

- **Endpoint**: `/api/v1/`
- **Method**: `GET`
- **Description**: Fetches a list of all peaks in the database with their relevant details (such as name, height, range, and whether they are open to the public).

```bash
# Example Request:
GET /api/v1/

# Example Response:
[
  {
    "Peak_id": 1,
    "Name": "Mount Everest",
    "Alias": "SGRM",
    "Height": 8848.86,
    "Peak_range": "Khumbu",
    "OpenToPublic": true
  },
  {
    "Peak_id": 2,
    "Name": "Kangchenjunga",
    "Alias": "KJN",
    "Height": 8586,
    "Peak_range": "Himalaya",
    "OpenToPublic": false
  }
]
```

### 2. Get a Specific Peak by Name

- **Endpoint**: `/api/v1/peak/{peakname}`
- **Method**: `GET`
- **Description**: Fetches detailed information about a specific peak by its name. Replace `{peakname}` with the actual name of the peak (e.g., `Mount Everest`).

```bash
# Example Request:
GET /api/v1/peak/Mount%20Everest

# Example Response:
{
  "Peak_id": 1,
  "Name": "Mount Everest",
  "Alias": "SGRM",
  "Height": 8848.86,
  "Peak_range": "Khumbu",
  "OpenToPublic": true
}
```

### 3. Search Peak by ID

- **Endpoint**: `/api/v1/search?id={num}`
- **Method**: `GET`
- **Description**: Retrieves peak information based on the unique `id` of the peak. Replace `{num}` with the numerical ID (e.g., `1` for Mount Everest).

```bash
# Example Request:
GET /api/v1/search?id=1

# Example Response:
{
  "Peak_id": 1,
  "Name": "Mount Everest",
  "Alias": "SGRM",
  "Height": 8848.86,
  "Peak_range": "Khumbu",
  "OpenToPublic": true
}
```

## Project Setup

This project is built with Go. To run the API locally, follow these steps:

```bash
1. Clone the repository:
   git clone https://github.com/RishavOsaurus/NepalMoutainsAPI.git

2. Navigate to the project directory:
   cd NepalMountainsAPI

3. Install the dependencies:
   go mod tidy

4. Run the API:
   go run main.go
```

The API will be available at `http://localhost:8080`.

## Data Source

The data used in this API was scraped from the official website of the [Nepal Himal Peak Profile](https://nepalhimalpeakprofile.org/peak-profile).
