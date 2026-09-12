# 🌄 Welcome to the Nepal Peaks API!

Are you ready to embark on an adventure through the towering peaks of Nepal? This **Go-based API** is your gateway to explore the majestic mountains, discover their heights, and learn about their significance. Let’s get started!

---

## 🚀 **Why Use This API?**

- **Discover Peaks:** Get detailed information about Nepal's famous mountains.
- **Easy Search:** Find specific peaks by name or unique ID.
- **Reliable Data:** All information is sourced from the official [Nepal Himal Peak Profile](https://nepalhimalpeakprofile.org/peak-profile) website.

---

## 📚 **Available Endpoints**

### 1. **Get All Peaks** 🌍
- **Endpoint:** `/api/v1/`
- **Method:** `GET`
- **What It Does:** Fetches a list of all peaks with their details.

**👉 Example Request:**
```http
GET /api/v1/
```

**🌟 Example Response:**
```json
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

### 2. **Get a Specific Peak by Name** 🏔️
- **Endpoint:** `/api/v1/peak/{peakname}`
- **Method:** `GET`
- **What It Does:** Fetches detailed info about a specific peak. Just replace `{peakname}`!

**👉 Example Request:**
```http
GET /api/v1/peak/Mount%20Everest
```

**🌟 Example Response:**
```json
{
  "Peak_id": 1,
  "Name": "Mount Everest",
  "Alias": "SGRM",
  "Height": 8848.86,
  "Peak_range": "Khumbu",
  "OpenToPublic": true
}
```

### 3. **Search Peak by ID** 🔍
- **Endpoint:** `/api/v1/search?id={num}`
- **Method:** `GET`
- **What It Does:** Retrieves peak information based on its unique ID.

**👉 Example Request:**
```http
GET /api/v1/search?id=1
```

**🌟 Example Response:**
```json
{
  "Peak_id": 1,
  "Name": "Mount Everest",
  "Alias": "SGRM",
  "Height": 8848.86,
  "Peak_range": "Khumbu",
  "OpenToPublic": true
}
```

---

## 🛠 **Getting Started: Set Up the API Locally**

Let’s bring this API to life on your machine! Follow these simple steps:

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/RishavOsaurus/NepalMoutainsAPI.git
   ```

2. **Navigate to the Project Directory:**
   ```bash
   cd NepalMountainsAPI
   ```

3. **Install the Dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the API:**
   ```bash
   go run main.go
   ```

🚀 Your API will be live at `http://localhost:8080`! Explore the endpoints and dive into the world of Nepal’s peaks.

---

## 🤝 **Contributions Welcome!**

We’d love your input! If you have ideas for enhancements or new features, feel free to open a pull request. Every contribution makes the API better!

### 📞 **Get in Touch**

Have questions or feedback? Reach out to [Rishav Osaurus](mailto:your-email@example.com). We’re here to help!

---

🌟 **Join Us in Exploring the Heights of Nepal’s Peaks!** 🌟

Dive into the adventure, learn about the stunning landscapes, and share your findings with fellow mountain enthusiasts! 🏔️✨
