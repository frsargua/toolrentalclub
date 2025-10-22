# Tool Rental Club

A full-stack application for tool rentals with React frontend, Golang backend, and Firebase authentication.

## 🚀 Quick Start

### Prerequisites

- Node.js 16+ and npm
- Go 1.21+
- Firebase project

### Setup

1. **Clone and navigate to the project**

```bash
cd toolrentalclub
```

2. **Set up the Frontend**

```bash
cd frontend
npm install
# Configure Firebase in src/config/firebase.ts
npm run dev
```

Frontend will run on `http://localhost:3000`

3. **Set up the Backend**

```bash
cd backend
go mod download
# Add your Firebase service account key as serviceAccountKey.json
# Create .env file from .env.example
go run cmd/api/main.go
```

Backend will run on `http://localhost:8080`

## 📁 Project Structure

```
toolrentalclub/
├── frontend/          # React + TypeScript + Vite
│   ├── src/
│   │   ├── components/    # React components
│   │   ├── services/      # API services
│   │   └── config/        # Firebase config
│   └── package.json
├── backend/           # Golang REST API
│   ├── cmd/api/          # Application entry point
│   ├── internal/         # Internal packages
│   │   ├── handlers/     # HTTP handlers
│   │   └── middleware/   # Middleware functions
│   └── go.mod
└── README.md
```

## 🔥 Firebase Setup

### 1. Create a Firebase Project

1. Go to [Firebase Console](https://console.firebase.google.com/)
2. Create a new project
3. Enable Authentication:
   - Email/Password provider
   - Google provider

### 2. Frontend Configuration

1. In Firebase Console, go to Project Settings
2. Add a web app and copy the config
3. Update `frontend/src/config/firebase.ts` with your config

### 3. Backend Configuration

1. Go to Project Settings > Service Accounts
2. Click "Generate New Private Key"
3. Save as `backend/serviceAccountKey.json`
4. Update `backend/.env` with the path

## 🎨 Features

### Frontend

- ✨ Modern, responsive UI with Bootstrap 5
- 🔐 Firebase authentication (Email/Password & Google)
- 📱 Mobile-friendly design with React Bootstrap components
- 🎯 TypeScript for type safety
- ⚡ Fast development with Vite

### Backend

- 🚀 RESTful API with Gorilla Mux
- 🔒 Firebase token verification
- 📝 Request logging middleware
- 🌐 CORS support
- 🏗️ Clean architecture

## 🔐 Authentication Flow

1. User signs up/logs in on frontend (Firebase)
2. Frontend receives Firebase ID token
3. Token sent to backend for verification
4. Backend verifies with Firebase Admin SDK
5. Protected routes require valid token in Authorization header

## 📡 API Endpoints

### Public

- `GET /api/health` - Health check

### Auth

- `POST /api/auth/verify` - Verify Firebase token

### Protected (require Bearer token)

- `GET /api/profile` - Get user profile

## 🛠️ Development

### Frontend Development

```bash
cd frontend
npm run dev      # Start dev server
npm run build    # Build for production
npm run preview  # Preview production build
```

### Backend Development

```bash
cd backend
go run cmd/api/main.go  # Start server
go test ./...           # Run tests
go build -o bin/server cmd/api/main.go  # Build binary
```

## 🚀 Deployment

### Frontend

Deploy to:

- Vercel
- Netlify
- Firebase Hosting
- Any static hosting service

```bash
cd frontend
npm run build
# Deploy the 'dist' folder
```

### Backend

Deploy to:

- Google Cloud Run
- Heroku
- DigitalOcean
- AWS

```bash
cd backend
go build -o server cmd/api/main.go
# Deploy the binary
```

## 📝 Environment Variables

### Frontend (.env)

```
VITE_FIREBASE_API_KEY=your_api_key
VITE_FIREBASE_AUTH_DOMAIN=your_domain
VITE_FIREBASE_PROJECT_ID=your_project_id
VITE_FIREBASE_STORAGE_BUCKET=your_bucket
VITE_FIREBASE_MESSAGING_SENDER_ID=your_sender_id
VITE_FIREBASE_APP_ID=your_app_id
```

### Backend (.env)

```
PORT=8080
FIREBASE_SERVICE_ACCOUNT=./serviceAccountKey.json
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## 📄 License

MIT

## 🆘 Support

For issues and questions:

- Frontend: Check `frontend/README.md`
- Backend: Check `backend/README.md`

## 🎯 Next Steps

- [ ] Add tool listing and management
- [ ] Implement rental booking system
- [ ] Add user profiles and history
- [ ] Set up payment integration
- [ ] Add search and filtering
- [ ] Implement reviews and ratings

Happy coding! 🎉
