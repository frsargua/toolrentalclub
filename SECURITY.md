# Security Guide - Tool Rental Club

## 🔒 Environment Variables & Secrets

This project uses environment variables to keep sensitive configuration secure and prevent accidental commits of credentials.

## Frontend Security (.env)

### What's Protected

The frontend `.env` file contains your Firebase configuration:

```bash
VITE_FIREBASE_API_KEY=your_key
VITE_FIREBASE_AUTH_DOMAIN=your_domain
# ... etc
```

### Important Notes

✅ **Protected:** The `.env` file is in `.gitignore` and **will not be committed**

✅ **Safe to expose:** Firebase API keys in frontend apps are designed to be public (they're sent to browsers). Security is enforced by:
- Firebase Security Rules
- App restrictions in Firebase Console
- Firebase Authentication

⚠️ **Still use .env:** Even though Firebase keys are "public", using `.env` keeps your codebase clean and makes it easy to:
- Switch between dev/staging/prod environments
- Share code without exposing your specific project
- Rotate credentials if needed

### Setup Instructions

1. **Your .env is already created** with your Firebase credentials
2. **Never commit `.env`** to version control (it's already in `.gitignore`)
3. **Share `.env.example`** with your team (without actual values)
4. **Each developer** creates their own `.env` from `.env.example`

## Backend Security

### Service Account Key

The backend uses a Firebase Admin SDK service account key:

```
backend/serviceAccountKey.json
```

⚠️ **CRITICAL:** This file contains **private credentials** and must NEVER be committed!

✅ **Protected:** Already in `.gitignore`

### Security Checklist

- [ ] `backend/serviceAccountKey.json` exists locally
- [ ] `backend/serviceAccountKey.json` is in `.gitignore`
- [ ] Never commit service account keys
- [ ] Restrict service account permissions in Firebase Console

## Git Security

### What's Already Protected

Your `.gitignore` files already protect:

**Frontend:**
```
.env
.env.local
.env.production
```

**Backend:**
```
serviceAccountKey.json
*-firebase-adminsdk-*.json
```

**Root:**
```
.env
.env.local
serviceAccountKey.json
*-firebase-adminsdk-*.json
```

### Verify Your Setup

Run this command to check what's being tracked:

```bash
git status --ignored
```

You should see `.env` and `serviceAccountKey.json` in the ignored files list.

## If You Accidentally Committed Secrets

If you've already committed sensitive files:

### 1. Remove from Git (keep local file)

```bash
# Remove .env from git but keep the file locally
git rm --cached frontend/.env

# Remove service account key from git but keep locally
git rm --cached backend/serviceAccountKey.json

# Commit the removal
git commit -m "Remove sensitive files from git"
```

### 2. Rotate Credentials

If credentials were pushed to a remote repository:

**Firebase Frontend:**
1. Go to Firebase Console
2. Restrict your API key to your domain only
3. Consider creating a new Firebase project if exposed publicly

**Firebase Backend (Service Account):**
1. Go to Firebase Console > Project Settings > Service Accounts
2. Delete the compromised service account
3. Create a new service account
4. Download new key
5. Update your local `serviceAccountKey.json`

### 3. Clean Git History (if needed)

If sensitive data is in your git history:

```bash
# Install BFG Repo Cleaner
brew install bfg

# Remove .env from all history
bfg --delete-files .env

# Clean up
git reflog expire --expire=now --all
git gc --prune=now --aggressive

# Force push (WARNING: this rewrites history)
git push origin --force --all
```

⚠️ **Warning:** Only do this if you haven't shared the repository with others.

## Firebase Security Rules

### Firestore Rules

Even with environment variables, you need proper security rules:

```javascript
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    // Require authentication for all reads/writes
    match /{document=**} {
      allow read, write: if request.auth != null;
    }
    
    // Example: Users can only read/write their own data
    match /users/{userId} {
      allow read, write: if request.auth != null && request.auth.uid == userId;
    }
  }
}
```

### Storage Rules

```javascript
rules_version = '2';
service firebase.storage {
  match /b/{bucket}/o {
    match /{allPaths=**} {
      allow read, write: if request.auth != null;
    }
  }
}
```

## API Key Restrictions (Recommended)

### Restrict Frontend API Key

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Select your Firebase project
3. Go to "APIs & Services" > "Credentials"
4. Find your API key
5. Click "Edit API key"
6. Under "Application restrictions":
   - Select "HTTP referrers (web sites)"
   - Add your domains:
     - `http://localhost:3000/*` (development)
     - `https://yourdomain.com/*` (production)

## Environment Best Practices

### Development

```bash
# frontend/.env
VITE_FIREBASE_PROJECT_ID=myapp-dev
```

### Staging

```bash
# frontend/.env.staging
VITE_FIREBASE_PROJECT_ID=myapp-staging
```

### Production

```bash
# frontend/.env.production
VITE_FIREBASE_PROJECT_ID=myapp-prod
```

Load different configs:

```bash
# Development
npm run dev

# Production build
npm run build  # Uses .env.production
```

## Team Collaboration

### Sharing Configuration

**DO:**
- ✅ Share `.env.example` (template without actual values)
- ✅ Document where to get credentials
- ✅ Use a password manager for team secrets (1Password, LastPass, etc.)

**DON'T:**
- ❌ Commit `.env` files
- ❌ Share credentials via email/Slack
- ❌ Push service account keys to git

### Onboarding New Developers

1. Clone the repository
2. Copy `.env.example` to `.env`
3. Get Firebase credentials from team (securely)
4. Get service account key from team (securely)
5. Start development!

## Production Deployment

### Hosting Platforms

Most hosting platforms (Vercel, Netlify, etc.) let you set environment variables in their dashboard:

**Vercel:**
1. Project Settings > Environment Variables
2. Add each `VITE_*` variable
3. Deploy

**Netlify:**
1. Site Settings > Build & Deploy > Environment
2. Add variables
3. Redeploy

### Backend Deployment

For backend (Go):
- Use platform secrets (Cloud Run secrets, AWS Secrets Manager, etc.)
- Mount service account key as a file or environment variable
- Never include service account key in Docker image

## Monitoring

### Check for Exposed Secrets

Use tools like:
- [git-secrets](https://github.com/awslabs/git-secrets)
- [truffleHog](https://github.com/trufflesecurity/trufflehog)
- GitHub secret scanning (automatic)

### Firebase Monitoring

Monitor your Firebase usage:
- Set up budget alerts in Firebase Console
- Monitor authentication attempts
- Review security rules regularly

## Questions?

- Unsure if something should be in `.env`? → Yes, add it
- Need to share secrets with team? → Use a secure password manager
- Credentials exposed? → Rotate immediately

Stay secure! 🔒

