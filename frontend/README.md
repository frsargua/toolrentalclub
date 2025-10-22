# Tool Rental Club - Frontend

A modern React application for Tool Rental Club with Firebase authentication and Bootstrap UI.

## Features

- ✨ Modern, clean UI built with React and TypeScript
- 🎨 Styled with Bootstrap 5 and React Bootstrap
- 🔐 Firebase Authentication (Email/Password & Google Sign-In)
- 🚀 Fast development with Vite
- 📱 Fully responsive design
- 🔗 Ready to integrate with Golang backend

## Prerequisites

- Node.js 16+ and npm
- A Firebase project (for authentication)

## Setup

### 1. Install Dependencies

```bash
npm install
```

### 2. Configure Firebase

1. Go to the [Firebase Console](https://console.firebase.google.com/)
2. Create a new project or select an existing one
3. Enable Authentication:
   - Go to Authentication > Sign-in method
   - Enable "Email/Password" provider
   - Enable "Google" provider
4. Get your Firebase config:

   - Go to Project Settings > General
   - Scroll down to "Your apps" and click the web icon (</>)
   - Copy the config object

5. Create a `.env` file in the frontend directory:

```bash
# Copy the example file
cp .env.example .env
```

6. Update `.env` with your Firebase configuration:

```bash
VITE_FIREBASE_API_KEY=your_api_key
VITE_FIREBASE_AUTH_DOMAIN=your_project.firebaseapp.com
VITE_FIREBASE_PROJECT_ID=your_project_id
VITE_FIREBASE_STORAGE_BUCKET=your_project.appspot.com
VITE_FIREBASE_MESSAGING_SENDER_ID=your_sender_id
VITE_FIREBASE_APP_ID=your_app_id
VITE_FIREBASE_MEASUREMENT_ID=your_measurement_id
```

**Important:** Never commit your `.env` file to Git. It's already in `.gitignore`.

### 3. Configure Backend Integration

The app is configured to proxy API calls to a Golang backend running on `http://localhost:8080`.

You can modify the backend URL in `vite.config.ts`:

```typescript
proxy: {
  "/api": {
    target: "http://localhost:8080",
    changeOrigin: true,
  },
};
```

## Development

Start the development server:

```bash
npm run dev
```

The app will be available at `http://localhost:3000`

## Building for Production

```bash
npm run build
```

The production-ready files will be in the `dist` directory.

## Project Structure

```
frontend/
├── src/
│   ├── components/
│   │   └── Login.tsx          # Login/Register component
│   ├── config/
│   │   └── firebase.ts        # Firebase configuration
│   ├── services/
│   │   └── api.ts             # API service functions
│   ├── App.tsx                # Main app component
│   ├── main.tsx               # Entry point
│   └── index.css              # Global styles
├── public/                    # Static assets
├── index.html                 # HTML template
├── vite.config.ts            # Vite configuration
└── package.json              # Dependencies and scripts
```

## Backend Integration

The frontend is set up to work with a Golang backend. Here's what you need to implement on the backend:

### 1. Token Verification Endpoint

```go
POST /api/auth/verify
Content-Type: application/json

{
  "token": "firebase-id-token"
}
```

This endpoint should verify the Firebase ID token and create/update the user session.

### 2. Example Go Backend Setup

Install the Firebase Admin SDK:

```bash
go get firebase.google.com/go/v4
```

Example verification code:

```go
// Verify Firebase token
func verifyToken(idToken string) (*auth.Token, error) {
    client, err := app.Auth(ctx)
    if err != nil {
        return nil, err
    }

    token, err := client.VerifyIDToken(ctx, idToken)
    if err != nil {
        return nil, err
    }

    return token, nil
}
```

## API Service

The `src/services/api.ts` file provides:

- `loginWithEmail(email, password)` - Sign in with email/password
- `registerWithEmail(email, password)` - Register new user
- `loginWithGoogle()` - Sign in with Google
- `logout()` - Sign out user
- `apiCall(endpoint, options)` - Generic API call with automatic token injection

## Styling with Bootstrap

This project uses **Bootstrap 5** and **React Bootstrap** for styling.

### Using Bootstrap Components

Import components from `react-bootstrap`:

```tsx
import { Button, Card, Form, Alert } from "react-bootstrap";

function MyComponent() {
  return (
    <Card>
      <Card.Body>
        <Form>
          <Form.Group>
            <Form.Label>Email</Form.Label>
            <Form.Control type="email" />
          </Form.Group>
          <Button variant="primary">Submit</Button>
        </Form>
      </Card.Body>
    </Card>
  );
}
```

### Bootstrap Utilities

Use Bootstrap utility classes for spacing, colors, etc:

```tsx
<div className="mt-4 p-3 bg-light rounded shadow">
  <h2 className="text-primary fw-bold">Hello</h2>
  <p className="text-muted">Description</p>
</div>
```

### Custom Styles

Add custom CSS in `src/index.css`:

```css
/* Custom styles */
.my-custom-class {
  /* Your styles */
}
```

### Customizing Bootstrap

To customize Bootstrap variables, you can create a custom SCSS file. First install sass:

```bash
npm install -D sass
```

Then create `src/custom.scss`:

```scss
// Override Bootstrap variables
$primary: #0ea5e9;
$font-family-base: "Inter", sans-serif;

// Import Bootstrap
@import "bootstrap/scss/bootstrap";
```

And import it in `main.tsx` instead of the CSS import.

## Creating Reusable Components

When you're ready to create custom reusable components, follow this structure:

```
src/
├── components/
│   ├── common/              # Reusable components
│   │   ├── Button.tsx
│   │   ├── Card.tsx
│   │   └── Input.tsx
│   ├── layout/              # Layout components
│   │   ├── Header.tsx
│   │   ├── Footer.tsx
│   │   └── Sidebar.tsx
│   └── features/            # Feature-specific components
│       ├── auth/
│       │   ├── Login.tsx
│       │   └── Register.tsx
│       └── tools/
│           ├── ToolList.tsx
│           └── ToolCard.tsx
```

## Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run lint` - Run ESLint

## Bootstrap Resources

- [Bootstrap Documentation](https://getbootstrap.com/)
- [React Bootstrap Documentation](https://react-bootstrap.github.io/)
- [Bootstrap Icons](https://icons.getbootstrap.com/)
- [Bootstrap Examples](https://getbootstrap.com/docs/5.3/examples/)

## License

MIT
