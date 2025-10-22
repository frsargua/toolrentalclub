import { useState, useEffect } from "react";
import { Container, Card, Row, Col, Spinner, Alert } from "react-bootstrap";
import { apiCall } from "../services/api";
import { auth } from "../config/firebase";

interface UserProfile {
  userId: string;
  email: string;
  message?: string;
}

function Profile() {
  const [profile, setProfile] = useState<UserProfile | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    fetchProfile();
  }, []);

  const fetchProfile = async () => {
    try {
      setLoading(true);
      const data = await apiCall("/profile");
      setProfile(data);
      setError("");
    } catch (err) {
      const errorMessage =
        err instanceof Error ? err.message : "Failed to load profile";
      setError(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <Container className="py-5">
        <div className="text-center">
          <Spinner animation="border" variant="primary" />
          <p className="mt-3 text-muted">Loading profile...</p>
        </div>
      </Container>
    );
  }

  if (error) {
    return (
      <Container className="py-5">
        <Alert variant="danger">
          <Alert.Heading>Error Loading Profile</Alert.Heading>
          <p>{error}</p>
        </Alert>
      </Container>
    );
  }

  const currentUser = auth.currentUser;

  return (
    <Container className="py-5">
      <Row className="justify-content-center">
        <Col xs={12} lg={8}>
          <Card className="shadow">
            <Card.Header className="bg-primary text-white">
              <h3 className="mb-0">User Profile</h3>
            </Card.Header>
            <Card.Body className="p-4">
              <div className="text-center mb-4">
                <div
                  className="d-inline-flex align-items-center justify-content-center bg-primary text-white rounded-circle mb-3"
                  style={{
                    width: "100px",
                    height: "100px",
                    fontSize: "2.5rem",
                  }}
                >
                  {currentUser?.email?.charAt(0).toUpperCase() || "U"}
                </div>
              </div>

              <Row className="mb-3">
                <Col md={4}>
                  <strong className="text-muted">User ID:</strong>
                </Col>
                <Col md={8}>
                  <code className="bg-light p-2 rounded d-inline-block">
                    {profile?.userId}
                  </code>
                </Col>
              </Row>

              <Row className="mb-3">
                <Col md={4}>
                  <strong className="text-muted">Email:</strong>
                </Col>
                <Col md={8}>{profile?.email || currentUser?.email}</Col>
              </Row>

              <Row className="mb-3">
                <Col md={4}>
                  <strong className="text-muted">Display Name:</strong>
                </Col>
                <Col md={8}>
                  {currentUser?.displayName || (
                    <span className="text-muted">Not set</span>
                  )}
                </Col>
              </Row>

              <Row className="mb-3">
                <Col md={4}>
                  <strong className="text-muted">Email Verified:</strong>
                </Col>
                <Col md={8}>
                  {currentUser?.emailVerified ? (
                    <span className="badge bg-success">Verified</span>
                  ) : (
                    <span className="badge bg-warning text-dark">
                      Not Verified
                    </span>
                  )}
                </Col>
              </Row>

              <Row className="mb-3">
                <Col md={4}>
                  <strong className="text-muted">Account Created:</strong>
                </Col>
                <Col md={8}>
                  {currentUser?.metadata?.creationTime
                    ? new Date(
                        currentUser.metadata.creationTime
                      ).toLocaleDateString()
                    : "Unknown"}
                </Col>
              </Row>

              <Row>
                <Col md={4}>
                  <strong className="text-muted">Last Sign In:</strong>
                </Col>
                <Col md={8}>
                  {currentUser?.metadata?.lastSignInTime
                    ? new Date(
                        currentUser.metadata.lastSignInTime
                      ).toLocaleString()
                    : "Unknown"}
                </Col>
              </Row>
            </Card.Body>
          </Card>

          {profile?.message && (
            <Alert variant="info" className="mt-4">
              <strong>Backend Message:</strong> {profile.message}
            </Alert>
          )}
        </Col>
      </Row>
    </Container>
  );
}

export default Profile;
