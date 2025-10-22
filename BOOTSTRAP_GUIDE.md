# Bootstrap Guide for Tool Rental Club

## Overview

This project uses **Bootstrap 5** and **React Bootstrap** for styling. This guide will help you work with Bootstrap components and utilities.

## Why Bootstrap?

- ✅ **Comprehensive component library** - Buttons, forms, modals, cards, and more
- ✅ **Responsive by default** - Mobile-first design
- ✅ **Easy to customize** - CSS variables and SCSS support
- ✅ **React Bootstrap** - Components designed for React
- ✅ **Large community** - Extensive documentation and examples

## Getting Started

### Importing Components

```tsx
// Import individual components
import { Button, Card, Form, Alert, Modal, Navbar } from "react-bootstrap";

// Use them in your component
function MyComponent() {
  return (
    <Card>
      <Card.Body>
        <Card.Title>Hello Bootstrap</Card.Title>
        <Button variant="primary">Click me</Button>
      </Card.Body>
    </Card>
  );
}
```

## Common Components

### 1. Buttons

```tsx
import { Button } from "react-bootstrap";

// Variants
<Button variant="primary">Primary</Button>
<Button variant="secondary">Secondary</Button>
<Button variant="success">Success</Button>
<Button variant="danger">Danger</Button>
<Button variant="warning">Warning</Button>
<Button variant="info">Info</Button>
<Button variant="light">Light</Button>
<Button variant="dark">Dark</Button>
<Button variant="link">Link</Button>

// Outline buttons
<Button variant="outline-primary">Outline Primary</Button>

// Sizes
<Button size="sm">Small</Button>
<Button size="lg">Large</Button>

// States
<Button disabled>Disabled</Button>
<Button active>Active</Button>
```

### 2. Forms

```tsx
import { Form, Button } from "react-bootstrap";

function LoginForm() {
  return (
    <Form>
      <Form.Group className="mb-3" controlId="email">
        <Form.Label>Email</Form.Label>
        <Form.Control type="email" placeholder="Enter email" />
        <Form.Text className="text-muted">
          We'll never share your email.
        </Form.Text>
      </Form.Group>

      <Form.Group className="mb-3" controlId="password">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" />
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Check type="checkbox" label="Remember me" />
      </Form.Group>

      <Button variant="primary" type="submit">
        Submit
      </Button>
    </Form>
  );
}
```

### 3. Cards

```tsx
import { Card, Button } from "react-bootstrap";

<Card style={{ width: "18rem" }}>
  <Card.Img variant="top" src="image.jpg" />
  <Card.Body>
    <Card.Title>Card Title</Card.Title>
    <Card.Text>Some quick example text to build on the card title.</Card.Text>
    <Button variant="primary">Go somewhere</Button>
  </Card.Body>
</Card>;
```

### 4. Alerts

```tsx
import { Alert } from "react-bootstrap";

<Alert variant="success">
  <Alert.Heading>Success!</Alert.Heading>
  <p>Your action was successful.</p>
</Alert>

// Dismissible alert
<Alert variant="danger" dismissible onClose={() => console.log("Closed")}>
  Error message here
</Alert>
```

### 5. Modals

```tsx
import { useState } from "react";
import { Modal, Button } from "react-bootstrap";

function MyModal() {
  const [show, setShow] = useState(false);

  return (
    <>
      <Button onClick={() => setShow(true)}>Open Modal</Button>

      <Modal show={show} onHide={() => setShow(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Modal Title</Modal.Title>
        </Modal.Header>
        <Modal.Body>Modal content goes here</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShow(false)}>
            Close
          </Button>
          <Button variant="primary" onClick={() => setShow(false)}>
            Save Changes
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
}
```

### 6. Navigation

```tsx
import { Navbar, Nav, Container } from "react-bootstrap";

<Navbar bg="light" expand="lg">
  <Container>
    <Navbar.Brand href="#home">Brand</Navbar.Brand>
    <Navbar.Toggle aria-controls="basic-navbar-nav" />
    <Navbar.Collapse id="basic-navbar-nav">
      <Nav className="me-auto">
        <Nav.Link href="#home">Home</Nav.Link>
        <Nav.Link href="#tools">Tools</Nav.Link>
        <Nav.Link href="#about">About</Nav.Link>
      </Nav>
    </Navbar.Collapse>
  </Container>
</Navbar>;
```

### 7. Tables

```tsx
import { Table } from "react-bootstrap";

<Table striped bordered hover>
  <thead>
    <tr>
      <th>#</th>
      <th>Tool Name</th>
      <th>Price</th>
      <th>Available</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>1</td>
      <td>Hammer</td>
      <td>$10/day</td>
      <td>Yes</td>
    </tr>
  </tbody>
</Table>;
```

## Bootstrap Utilities

Bootstrap provides many utility classes for common styling needs:

### Spacing

```tsx
// Margin
<div className="m-3">Margin on all sides</div>
<div className="mt-3">Margin top</div>
<div className="mb-3">Margin bottom</div>
<div className="ms-3">Margin start (left)</div>
<div className="me-3">Margin end (right)</div>
<div className="mx-3">Margin horizontal</div>
<div className="my-3">Margin vertical</div>

// Padding (same pattern with 'p')
<div className="p-4">Padding on all sides</div>
<div className="pt-4">Padding top</div>

// Sizes: 0, 1, 2, 3, 4, 5, auto
```

### Colors

```tsx
// Text colors
<p className="text-primary">Primary text</p>
<p className="text-success">Success text</p>
<p className="text-danger">Danger text</p>
<p className="text-warning">Warning text</p>
<p className="text-muted">Muted text</p>

// Background colors
<div className="bg-primary text-white">Primary background</div>
<div className="bg-light">Light background</div>
<div className="bg-dark text-white">Dark background</div>
```

### Display & Layout

```tsx
// Display
<div className="d-none">Hidden</div>
<div className="d-block">Block</div>
<div className="d-flex">Flexbox</div>
<div className="d-inline-flex">Inline flex</div>

// Flexbox utilities
<div className="d-flex justify-content-center">Centered</div>
<div className="d-flex align-items-center">Vertically centered</div>
<div className="d-flex justify-content-between">Space between</div>

// Responsive display
<div className="d-none d-md-block">Hidden on mobile, visible on tablet+</div>
```

### Sizing

```tsx
// Width
<div className="w-25">Width 25%</div>
<div className="w-50">Width 50%</div>
<div className="w-75">Width 75%</div>
<div className="w-100">Width 100%</div>

// Height
<div className="h-100">Height 100%</div>
<div className="vh-100">Height 100vh</div>

// Max width
<div className="mw-100">Max width 100%</div>
```

### Text

```tsx
// Alignment
<p className="text-start">Left aligned</p>
<p className="text-center">Center aligned</p>
<p className="text-end">Right aligned</p>

// Font weight
<p className="fw-bold">Bold text</p>
<p className="fw-normal">Normal weight</p>
<p className="fw-light">Light weight</p>

// Font size
<p className="fs-1">Largest</p>
<p className="fs-6">Smallest</p>

// Text transform
<p className="text-uppercase">UPPERCASE</p>
<p className="text-lowercase">lowercase</p>
<p className="text-capitalize">Capitalized</p>
```

### Borders & Rounded Corners

```tsx
// Borders
<div className="border">Border on all sides</div>
<div className="border-top">Top border only</div>
<div className="border-0">No border</div>

// Border colors
<div className="border border-primary">Primary border</div>

// Rounded corners
<div className="rounded">Rounded corners</div>
<div className="rounded-circle">Circle</div>
<div className="rounded-pill">Pill shape</div>
```

### Shadows

```tsx
<div className="shadow-sm">Small shadow</div>
<div className="shadow">Regular shadow</div>
<div className="shadow-lg">Large shadow</div>
```

## Layout Components

### Container

```tsx
import { Container, Row, Col } from "react-bootstrap";

// Responsive container
<Container>
  <Row>
    <Col>Column 1</Col>
    <Col>Column 2</Col>
    <Col>Column 3</Col>
  </Row>
</Container>

// Full-width container
<Container fluid>
  Content spans full width
</Container>

// Responsive columns
<Container>
  <Row>
    <Col xs={12} md={6} lg={4}>
      12 cols on mobile, 6 on tablet, 4 on desktop
    </Col>
  </Row>
</Container>
```

## Responsive Breakpoints

Bootstrap breakpoints:

- `xs` - Extra small (< 576px) - Default
- `sm` - Small (≥ 576px)
- `md` - Medium (≥ 768px)
- `lg` - Large (≥ 992px)
- `xl` - Extra large (≥ 1200px)
- `xxl` - Extra extra large (≥ 1400px)

Example usage:

```tsx
// Responsive spacing
<div className="mt-2 mt-md-4 mt-lg-5">
  Different margins on different screens
</div>

// Responsive display
<div className="d-none d-md-block">
  Hidden on mobile, visible on tablet and up
</div>
```

## Customizing Bootstrap

### Method 1: CSS Variables (Easiest)

Add to `frontend/src/index.css`:

```css
:root {
  --bs-primary: #0ea5e9;
  --bs-primary-rgb: 14, 165, 233;
  --bs-success: #10b981;
  --bs-danger: #ef4444;
  --bs-font-sans-serif: "Inter", system-ui, -apple-system, sans-serif;
}
```

### Method 2: Custom SCSS (More Control)

1. Install Sass:

```bash
npm install -D sass
```

2. Create `frontend/src/custom.scss`:

```scss
// 1. Include Bootstrap functions first
@import "bootstrap/scss/functions";

// 2. Override default variables
$primary: #0ea5e9;
$secondary: #64748b;
$success: #10b981;
$danger: #ef4444;
$font-family-base: "Inter", sans-serif;

// 3. Include Bootstrap
@import "bootstrap/scss/variables";
@import "bootstrap/scss/mixins";
@import "bootstrap/scss/root";
@import "bootstrap/scss/reboot";
@import "bootstrap/scss/type";
@import "bootstrap/scss/images";
@import "bootstrap/scss/containers";
@import "bootstrap/scss/grid";
@import "bootstrap/scss/tables";
@import "bootstrap/scss/forms";
@import "bootstrap/scss/buttons";
@import "bootstrap/scss/transitions";
@import "bootstrap/scss/dropdown";
@import "bootstrap/scss/button-group";
@import "bootstrap/scss/nav";
@import "bootstrap/scss/navbar";
@import "bootstrap/scss/card";
@import "bootstrap/scss/accordion";
@import "bootstrap/scss/breadcrumb";
@import "bootstrap/scss/pagination";
@import "bootstrap/scss/badge";
@import "bootstrap/scss/alert";
@import "bootstrap/scss/progress";
@import "bootstrap/scss/list-group";
@import "bootstrap/scss/close";
@import "bootstrap/scss/toasts";
@import "bootstrap/scss/modal";
@import "bootstrap/scss/tooltip";
@import "bootstrap/scss/popover";
@import "bootstrap/scss/carousel";
@import "bootstrap/scss/spinners";
@import "bootstrap/scss/offcanvas";
@import "bootstrap/scss/placeholders";
@import "bootstrap/scss/helpers";
@import "bootstrap/scss/utilities/api";
```

3. Update `frontend/src/main.tsx`:

```tsx
import "./custom.scss"; // Instead of index.css
```

## Tips for Creating Reusable Components

When you're ready to create custom components:

1. **Use Bootstrap as a foundation:**

```tsx
// Custom Button component
import { Button, ButtonProps } from "react-bootstrap";

interface CustomButtonProps extends ButtonProps {
  icon?: React.ReactNode;
}

export function CustomButton({ icon, children, ...props }: CustomButtonProps) {
  return (
    <Button {...props}>
      {icon && <span className="me-2">{icon}</span>}
      {children}
    </Button>
  );
}
```

2. **Combine Bootstrap utilities:**

```tsx
// Custom Card component
import { Card } from "react-bootstrap";

export function ToolCard({ tool }) {
  return (
    <Card className="h-100 shadow-sm hover-shadow">
      <Card.Img variant="top" src={tool.image} />
      <Card.Body>
        <Card.Title className="text-truncate">{tool.name}</Card.Title>
        <Card.Text className="text-muted">{tool.description}</Card.Text>
      </Card.Body>
      <Card.Footer className="bg-transparent border-0">
        <div className="d-flex justify-content-between align-items-center">
          <span className="fw-bold text-primary">${tool.price}/day</span>
          <Button size="sm" variant="primary">
            Rent
          </Button>
        </div>
      </Card.Footer>
    </Card>
  );
}
```

## Resources

- [React Bootstrap Documentation](https://react-bootstrap.github.io/)
- [Bootstrap Documentation](https://getbootstrap.com/docs/5.3/)
- [Bootstrap Icons](https://icons.getbootstrap.com/)
- [Bootstrap Examples](https://getbootstrap.com/docs/5.3/examples/)
- [Bootstrap Themes](https://themes.getbootstrap.com/)

## Next Steps

1. Explore the React Bootstrap components in the [official documentation](https://react-bootstrap.github.io/components/alerts)
2. Check out [Bootstrap examples](https://getbootstrap.com/docs/5.3/examples/) for inspiration
3. Start building your custom components when needed
4. Consider adding [Bootstrap Icons](https://icons.getbootstrap.com/) for additional icons

Happy coding! 🎨
