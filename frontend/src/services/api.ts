import {
  signInWithEmailAndPassword,
  createUserWithEmailAndPassword,
  signInWithPopup,
  signOut,
  UserCredential,
} from "firebase/auth";
import { auth, googleProvider } from "../config/firebase";

// Firebase Authentication Functions
export const loginWithEmail = async (
  email: string,
  password: string
): Promise<UserCredential> => {
  try {
    const userCredential = await signInWithEmailAndPassword(
      auth,
      email,
      password
    );

    // Optionally send the Firebase token to your Go backend
    const idToken = await userCredential.user.getIdToken();
    await sendTokenToBackend(idToken);

    return userCredential;
  } catch (error: any) {
    console.error("Login error:", error);
    throw new Error(error.message);
  }
};

export const registerWithEmail = async (
  email: string,
  password: string
): Promise<UserCredential> => {
  try {
    const userCredential = await createUserWithEmailAndPassword(
      auth,
      email,
      password
    );

    // Optionally send the Firebase token to your Go backend
    const idToken = await userCredential.user.getIdToken();
    await sendTokenToBackend(idToken);

    return userCredential;
  } catch (error: any) {
    console.error("Registration error:", error);
    throw new Error(error.message);
  }
};

export const loginWithGoogle = async (): Promise<UserCredential> => {
  try {
    const userCredential = await signInWithPopup(auth, googleProvider);

    // Optionally send the Firebase token to your Go backend
    const idToken = await userCredential.user.getIdToken();
    await sendTokenToBackend(idToken);

    return userCredential;
  } catch (error: any) {
    console.error("Google login error:", error);
    throw new Error(error.message);
  }
};

export const logout = async (): Promise<void> => {
  try {
    await signOut(auth);
  } catch (error: any) {
    console.error("Logout error:", error);
    throw new Error(error.message);
  }
};

// Backend API Functions
const API_BASE_URL = "/api";

export const sendTokenToBackend = async (idToken: string): Promise<void> => {
  try {
    const response = await fetch(`${API_BASE_URL}/auth/verify`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ token: idToken }),
    });

    if (!response.ok) {
      throw new Error("Failed to verify token with backend");
    }

    const data = await response.json();
    console.log("Token verified:", data);
  } catch (error: any) {
    console.error("Backend verification error:", error);
    // Don't throw - allow login even if backend verification fails
    // throw new Error(error.message)
  }
};

export const verifyTokenWithBackend = async (): Promise<boolean> => {
  try {
    const user = auth.currentUser;
    if (!user) {
      return false;
    }

    const idToken = await user.getIdToken(true); // force refresh token
    const response = await fetch(`${API_BASE_URL}/auth/verify`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ token: idToken }),
    });

    return response.ok;
  } catch (error: any) {
    console.error("Token verification failed:", error);
    return false;
  }
};

// Generic API call helper
export const apiCall = async (
  endpoint: string,
  options: RequestInit = {}
): Promise<any> => {
  try {
    const user = auth.currentUser;
    const idToken = user ? await user.getIdToken() : null;

    const headers: Record<string, string> = {
      "Content-Type": "application/json",
      ...((options.headers as Record<string, string>) || {}),
    };

    if (idToken) {
      headers["Authorization"] = `Bearer ${idToken}`;
    }

    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      ...options,
      headers,
    });

    if (!response.ok) {
      throw new Error(`API call failed: ${response.statusText}`);
    }

    return await response.json();
  } catch (error: any) {
    console.error("API call error:", error);
    throw new Error(error.message);
  }
};
