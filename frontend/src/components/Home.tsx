import { Container, Card, Row, Col, Button } from "react-bootstrap";
import { Link } from "react-router-dom";

function Home() {
  return (
    <Container className="py-5">
      <Row className="justify-content-center">
        <Col xs={12} lg={10}>
          <Card className="shadow-lg border-0 mb-4">
            <Card.Body className="p-5">
              <h1 className="display-4 fw-bold text-primary mb-4">
                Welcome to Tool Rental Club!
              </h1>
              <p className="lead text-muted mb-4">
                You've successfully logged in. Start exploring our tools and
                manage your rentals.
              </p>

              <Row className="g-4 mt-3">
                <Col md={6}>
                  <Card className="h-100 border-primary">
                    <Card.Body>
                      <div className="text-primary mb-3">
                        <svg
                          width="48"
                          height="48"
                          fill="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 3c1.66 0 3 1.34 3 3s-1.34 3-3 3-3-1.34-3-3 1.34-3 3-3zm0 14.2c-2.5 0-4.71-1.28-6-3.22.03-1.99 4-3.08 6-3.08 1.99 0 5.97 1.09 6 3.08-1.29 1.94-3.5 3.22-6 3.22z" />
                        </svg>
                      </div>
                      <Card.Title>Your Profile</Card.Title>
                      <Card.Text className="text-muted">
                        View and manage your account information and settings.
                      </Card.Text>
                      <Link to="/profile">
                        <Button variant="outline-primary">View Profile</Button>
                      </Link>
                    </Card.Body>
                  </Card>
                </Col>

                <Col md={6}>
                  <Card className="h-100 border-success">
                    <Card.Body>
                      <div className="text-success mb-3">
                        <svg
                          width="48"
                          height="48"
                          fill="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
                        </svg>
                      </div>
                      <Card.Title>Browse Tools</Card.Title>
                      <Card.Text className="text-muted">
                        Explore our collection of tools available for rent.
                      </Card.Text>
                      <Button variant="outline-success" disabled>
                        Coming Soon
                      </Button>
                    </Card.Body>
                  </Card>
                </Col>
              </Row>
            </Card.Body>
          </Card>

          <Card className="bg-light border-0">
            <Card.Body>
              <h5 className="mb-3">Getting Started</h5>
              <ul className="mb-0">
                <li>Complete your profile information</li>
                <li>Browse available tools</li>
                <li>Make your first rental</li>
                <li>Leave reviews and ratings</li>
              </ul>
            </Card.Body>
          </Card>
        </Col>
      </Row>
    </Container>
  );
}

export default Home;
