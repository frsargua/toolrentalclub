# Quick Start Guide - Tool Rental Club

## 🎯 What You Have

A complete authentication-ready application with:

- ✅ React frontend with beautiful Bootstrap UI
- ✅ Golang backend with Firebase integration
- ✅ Bootstrap 5 and React Bootstrap for styling
- ✅ All dependencies installed
- ✅ Ready to run locally

## 🚀 Get Started in 3 Steps

### Step 1: Set Up Firebase (5 minutes)

1. **Create Firebase Project**

   - Go to https://console.firebase.google.com/
   - Click "Add project"
   - Follow the setup wizard

2. **Enable Authentication**

   - In Firebase Console, go to "Authentication" → "Sign-in method"
   - Enable "Email/Password"
   - Enable "Google" sign-in

3. **Get Frontend Config**

   - Go to Project Settings (⚙️ icon)
   - Scroll to "Your apps" → Click Web icon (</>)
   - Copy the config values
   - Open `frontend/.env` file
   - Replace the placeholder values:

   ```bash
   VITE_FIREBASE_API_KEY=your_actual_api_key
   VITE_FIREBASE_AUTH_DOMAIN=your-project.firebaseapp.com
   VITE_FIREBASE_PROJECT_ID=your-project-id
   VITE_FIREBASE_STORAGE_BUCKET=your-project.appspot.com
   VITE_FIREBASE_MESSAGING_SENDER_ID=123456789
   VITE_FIREBASE_APP_ID=your_app_id
   VITE_FIREBASE_MEASUREMENT_ID=your_measurement_id
   ```

   **Note:** If `.env` doesn't exist, copy `.env.example` to `.env` first

4. **Get Backend Service Account**
   - In Firebase Console, go to Project Settings → Service Accounts
   - Click "Generate new private key"
   - Save the file as `backend/serviceAccountKey.json`

### Step 2: Run the Backend

```bash
cd backend
go run cmd/api/main.go
```

You should see:

```
Server starting on port 8080
```

**Note:** The backend will show a warning about Firebase not being configured until you add the service account key. The app will still work for testing!

### Step 3: Run the Frontend

Open a new terminal:

```bash
cd frontend
npm run dev
```

You should see:

```
  VITE v5.0.8  ready in 500 ms

  ➜  Local:   http://localhost:3000/
```

## 🎉 Test It Out!

1. Open http://localhost:3000 in your browser
2. You'll see the beautiful login page
3. Click "Don't have an account? Sign up"
4. Create a test account with email/password
5. Or sign in with Google!

## 📁 Project Structure

```
toolrentalclub/
├── frontend/              # React app
│   ├── src/
│   │   ├── components/
│   │   │   └── Login.tsx     # 🎨 Beautiful login UI
│   │   ├── config/
│   │   │   └── firebase.ts   # 🔥 Configure this!
│   │   ├── services/
│   │   │   └── api.ts        # API integration
│   │   └── App.tsx
│   └── package.json
│
├── backend/               # Go API
│   ├── cmd/api/
│   │   └── main.go          # Entry point
│   ├── internal/
│   │   ├── handlers/        # API endpoints
│   │   └── middleware/      # Auth, CORS, logging
│   └── serviceAccountKey.json  # 🔑 Add this!
│
└── README.md
```

## 🔧 Configuration Files

### Frontend: `frontend/src/config/firebase.ts`

```typescript
const firebaseConfig = {
  apiKey: "...", // From Firebase Console
  authDomain: "...",
  projectId: "...",
  storageBucket: "...",
  messagingSenderId: "...",
  appId: "...",
};
```

### Backend: `backend/serviceAccountKey.json`

Download from Firebase Console → Project Settings → Service Accounts

## 📡 How It Works

1. **User logs in** on frontend (React + Firebase)
2. **Firebase returns** an ID token
3. **Frontend sends token** to backend at `/api/auth/verify`
4. **Backend verifies** with Firebase Admin SDK
5. **Protected routes** require Bearer token in headers

## 🛠️ Development Commands

### Frontend

```bash
cd frontend
npm run dev        # Start dev server (port 3000)
npm run build      # Build for production
npm run preview    # Preview production build
```

### Backend

```bash
cd backend
go run cmd/api/main.go          # Start server (port 8080)
go build -o bin/server cmd/api/main.go  # Build binary
./bin/server                    # Run binary
```

## 🔌 API Endpoints

### Public Routes

- `GET /api/health` - Health check

### Auth Routes

- `POST /api/auth/verify` - Verify Firebase token
  ```json
  {
    "token": "firebase-id-token"
  }
  ```

### Protected Routes (require Bearer token)

- `GET /api/profile` - Get user profile

## 🐛 Troubleshooting

### Frontend Issues

**Problem:** Firebase errors

- ✅ Check `firebase.ts` has correct config
- ✅ Enable authentication in Firebase Console
- ✅ Check browser console for errors

**Problem:** Can't connect to backend

- ✅ Make sure backend is running on port 8080
- ✅ Check `vite.config.ts` proxy settings

### Backend Issues

**Problem:** "Firebase not configured" warning

- ✅ Add `serviceAccountKey.json` to backend folder
- ✅ Check file path in `.env` or use default location

**Problem:** Port 8080 already in use

- ✅ Change port: `export PORT=8081` then run server
- ✅ Update frontend `vite.config.ts` proxy

## 🎨 Customization

### Change Colors

Bootstrap uses CSS variables. You can customize colors by:

1. **Using Bootstrap utilities:**

   ```tsx
   <Button variant="primary">Click me</Button>
   <div className="bg-success text-white">...</div>
   ```

2. **Custom CSS (add to `frontend/src/index.css`):**

   ```css
   :root {
     --bs-primary: #0ea5e9;
     --bs-primary-rgb: 14, 165, 233;
   }
   ```

3. **For advanced customization, use SCSS:**
   ```bash
   cd frontend
   npm install -D sass
   ```
   Create `src/custom.scss`:
   ```scss
   $primary: #0ea5e9;
   @import "bootstrap/scss/bootstrap";
   ```

### Add New Routes

1. Create handler in `backend/internal/handlers/`
2. Register route in `backend/cmd/api/main.go`
3. Create API function in `frontend/src/services/api.ts`

## 🚀 Next Steps

Now that you have authentication working, you can:

- [ ] Add user profile page
- [ ] Create tool listing pages
- [ ] Implement rental booking system
- [ ] Add payment integration
- [ ] Set up database (Firestore is already configured!)

## 💡 Tips

1. **Use the auth context**: The authenticated user's ID is available in `r.Context().Value("userID")` in Go
2. **Check the READMEs**: Both `frontend/README.md` and `backend/README.md` have detailed docs
3. **Firebase Firestore**: Already set up! Use it for your database
4. **Environment variables**: Never commit `.env` or `serviceAccountKey.json`

## 📚 Learn More

- [React Documentation](https://react.dev/)
- [Vite Guide](https://vitejs.dev/guide/)
- [Firebase Auth Docs](https://firebase.google.com/docs/auth)
- [Go Firebase Admin SDK](https://firebase.google.com/docs/admin/setup)
- [Gorilla Mux](https://github.com/gorilla/mux)

## 🆘 Need Help?

Check the detailed READMEs:

- Main project: `README.md`
- Frontend: `frontend/README.md`
- Backend: `backend/README.md`

---

Happy coding! 🎉 If you have any questions, check the documentation or the code comments.
