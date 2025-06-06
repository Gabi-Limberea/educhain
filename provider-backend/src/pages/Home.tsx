import React, { FC, useState } from 'react';
import { Container, Row, Col, Card, Button, Offcanvas } from 'react-bootstrap';
import 'bootstrap-icons/font/bootstrap-icons.css';

const Home: FC = () => {
	const [showMenu, setShowMenu] = useState(false);

	const utilLinks = [
		{ href: '/account', label: 'Account', icon: 'bi-person' },
		{ href: '/settings', label: 'Settings', icon: 'bi-gear' },
		{ href: '/help', label: 'Help', icon: 'bi-question-circle' },
	];

	const scrollToSection = (id: string) => {
		const el = document.getElementById(id);
		if (el) el.scrollIntoView({ behavior: 'smooth', block: 'start' });
		setShowMenu(false);
	};

	return (
		<div className="bg-white">
			{/* Collapsible Sticky Sidebar Navigation */}
			<div 
				className="position-fixed start-0 top-50 translate-middle-y d-none d-lg-block sidebar-container" 
				style={{ 
					zIndex: 1000,
					transition: 'all 0.3s ease-in-out'
				}}
			>
				<div 
					className="bg-white shadow rounded-end p-2 sidebar-content"
					style={{ 
						borderRadius: '0 10px 10px 0',
						transition: 'all 0.3s ease-in-out',
						overflow: 'hidden'
					}}
				>
					{/* Collapsed state - just the icon */}
					<div className="sidebar-collapsed">
						<div className="d-flex align-items-center justify-content-center" style={{ width: '40px', height: '40px' }}>
							<i className="bi bi-list text-secondary" style={{ fontSize: '1.2rem' }}></i>
						</div>
					</div>
					
					{/* Expanded state - full navigation */}
					<div className="sidebar-expanded">
						<div className="d-flex align-items-center mb-3 px-2">
							<i className="bi bi-list text-secondary me-2" style={{ fontSize: '1.2rem' }}></i>
							<small className="text-muted fw-bold">NAVIGATION</small>
						</div>
						<div className="d-flex flex-column gap-2">
							<button
								onClick={() => scrollToSection('overview')}
								className="btn btn-sm btn-outline-secondary text-start px-3 py-2"
								style={{ fontSize: '0.875rem', minWidth: '140px' }}
							>
								<i className="bi bi-house me-2"></i>
								Overview
							</button>
							<button
								onClick={() => scrollToSection('features')}
								className="btn btn-sm btn-outline-secondary text-start px-3 py-2"
								style={{ fontSize: '0.875rem', minWidth: '140px' }}
							>
								<i className="bi bi-star me-2"></i>
								Features
							</button>
							<button
								onClick={() => scrollToSection('trusted')}
								className="btn btn-sm btn-outline-secondary text-start px-3 py-2"
								style={{ fontSize: '0.875rem', minWidth: '140px' }}
							>
								<i className="bi bi-shield-check me-2"></i>
								Made for
							</button>
						</div>
					</div>
				</div>
			</div>

			{/* Add CSS styles for the collapsible sidebar */}
			<style dangerouslySetInnerHTML={{
				__html: `
					.sidebar-container {
						width: 44px;
					}
					
					.sidebar-container:hover {
						width: auto;
					}
					
					.sidebar-collapsed {
						display: block;
					}
					
					.sidebar-expanded {
						display: none;
					}
					
					.sidebar-container:hover .sidebar-collapsed {
						display: none;
					}
					
					.sidebar-container:hover .sidebar-expanded {
						display: block;
					}
					
					.sidebar-content {
						min-height: 44px;
					}
					
					.sidebar-container:hover .sidebar-content {
						min-height: auto;
					}
				`
			}} />

			{/* Mobile Menu Button */}
			<div className="position-fixed top-0 end-0 p-3 d-lg-none" style={{ zIndex: 1000 }}>
				<Button 
					variant="outline-secondary" 
					onClick={() => setShowMenu(true)}
				>
					<i className="bi bi-list"></i>
				</Button>
			</div>

			{/* Hero Section */}
			<section id="overview" className="py-5 bg-light">
				<Container>
					<Row className="align-items-center min-vh-75">
						<Col lg={6}>
							<h1 className="display-4 fw-bold text-dark mb-4">
								The Blockchain Credentialing Platform
							</h1>
							<p className="lead text-muted mb-4">
								The leading platform to design, issue, and manage blockchain-verified digital credentials. 
								Secure, verifiable, and trusted by educational institutions worldwide.
							</p>
							<div className="d-flex gap-3 flex-wrap">
								<Button variant="primary" size="lg" href="/signup">
									Get Started Free
								</Button>
								<Button variant="outline-secondary" size="lg" href="/help">
									See How It Works
								</Button>
							</div>
						</Col>
						<Col lg={6} className="text-center">
							<div className="bg-white rounded shadow-lg p-4 d-inline-block">
								<i className="bi bi-award" style={{ fontSize: '8rem', color: '#6c757d' }}></i>
								<div className="mt-3">
									<small className="text-muted">Blockchain Verified</small>
									<div className="d-flex justify-content-center align-items-center mt-2">
									</div>
								</div>
							</div>
						</Col>
					</Row>
				</Container>
			</section>

			{/* Made for Section */}
			<section id="trusted" className="py-5 bg-white">
				<Container>
					<div className="text-center mb-5">
						<h3 className="fw-bold text-dark">Made for</h3>
					</div>
					<Row className="align-items-center justify-content-center text-center">
						<Col md={2} className="mb-3">
							<div className="text-muted">
								<i className="bi bi-bank" style={{ fontSize: '2rem' }}></i>
								<div className="mt-2 small">Educational Institutions</div>
							</div>
						</Col>
						<Col md={2} className="mb-3">
							<div className="text-muted">
								<i className="bi bi-building" style={{ fontSize: '2rem' }}></i>
								<div className="mt-2 small">Universities</div>
							</div>
						</Col>
						<Col md={2} className="mb-3">
							<div className="text-muted">
								<i className="bi bi-mortarboard" style={{ fontSize: '2rem' }}></i>
								<div className="mt-2 small">Bootcamps</div>
							</div>
						</Col>
					</Row>
				</Container>
			</section>

			{/* Features Section */}
			<section id="features" className="py-5 bg-light">
				<Container>
					<Row className="g-5">
						<Col lg={6}>
							<Card className="border-0 shadow-sm h-100">
								<Card.Body className="p-4">
									<div className="d-flex align-items-center mb-3">
										<div className="bg-primary bg-opacity-10 rounded p-2 me-3">
											<i className="bi bi-palette text-primary" style={{ fontSize: '1.5rem' }}></i>
										</div>
										<h4 className="fw-bold text-dark mb-0">Design and Issue Credentials</h4>
									</div>
									<p className="text-muted">
										Get access to the entire credentialing value-chain. Design your verifiable credentials 
										with your branding and style, issue them to your users, and then manage their lifecycle.
									</p>
								</Card.Body>
							</Card>
						</Col>
						<Col lg={6}>
							<Card className="border-0 shadow-sm h-100">
								<Card.Body className="p-4">
									<div className="d-flex align-items-center mb-3">
										<div className="bg-success bg-opacity-10 rounded p-2 me-3">
											<i className="bi bi-shield-check text-success" style={{ fontSize: '1.5rem' }}></i>
										</div>
										<h4 className="fw-bold text-dark mb-0">Instant Blockchain Verification</h4>
									</div>
									<p className="text-muted">
										Ensure your awards are trusted by third-parties. Our certificates are automatically 
										uploaded to the blockchain for secure and instant verification by employers and auditors.
									</p>
								</Card.Body>
							</Card>
						</Col>
						<Col lg={6}>
							<Card className="border-0 shadow-sm h-100">
								<Card.Body className="p-4">
									<div className="d-flex align-items-center mb-3">
										<div className="bg-info bg-opacity-10 rounded p-2 me-3">
											<i className="bi bi-wallet2 text-info" style={{ fontSize: '1.5rem' }}></i>
										</div>
										<h4 className="fw-bold text-dark mb-0">No Hidden Costs</h4>
									</div>
									<p className="text-muted">
										You don't incur any costs when uploading data to the blockchain. Issue unlimited 
										credentials to your users without being surprised by additional fees.
									</p>
								</Card.Body>
							</Card>
						</Col>
						<Col lg={6}>
							<Card className="border-0 shadow-sm h-100">
								<Card.Body className="p-4">
									<div className="d-flex align-items-center mb-3">
										<div className="bg-warning bg-opacity-10 rounded p-2 me-3">
											<i className="bi bi-leaf text-warning" style={{ fontSize: '1.5rem' }}></i>
										</div>
										<h4 className="fw-bold text-dark mb-0">Eco-Friendly Blockchain</h4>
									</div>
									<p className="text-muted">
										We're leading the way in sustainable blockchain technology, so you can get all the 
										benefits of blockchain verification without sacrificing environmental responsibility.
									</p>
								</Card.Body>
							</Card>
						</Col>
					</Row>
				</Container>
			</section>

			{/* Quick Actions Section */}
			<section className="py-5 bg-white">
				<Container>
					<div className="text-center mb-5">
						<h3 className="fw-bold text-dark">Get Started Today</h3>
						<p className="text-muted">Choose your path to blockchain-verified credentials</p>
					</div>
					<Row className="justify-content-center">
						<Col md={4} className="mb-4">
							<Card className="border-0 shadow-sm text-center h-100">
								<Card.Body className="p-4">
									<i className="bi bi-person-plus text-primary mb-3" style={{ fontSize: '3rem' }}></i>
									<h5 className="fw-bold">Register Students</h5>
									<p className="text-muted mb-4">Add new students to the platform and manage their academic journey</p>
									<Button variant="outline-primary" href="/registerStudents" className="w-100">
										Get Started
									</Button>
								</Card.Body>
							</Card>
						</Col>
						<Col md={4} className="mb-4">
							<Card className="border-0 shadow-sm text-center h-100">
								<Card.Body className="p-4">
									<i className="bi bi-award text-success mb-3" style={{ fontSize: '3rem' }}></i>
									<h5 className="fw-bold">Register Certificates</h5>
									<p className="text-muted mb-4">Create blockchain-verified digital certificates for your contests</p>
									<Button variant="outline-success" href="/registerCertification" className="w-100">
										Issue Now
									</Button>
								</Card.Body>
							</Card>
						</Col>
					</Row>
				</Container>
			</section>

			{/* Mobile Menu */}
			<Offcanvas show={showMenu} onHide={() => setShowMenu(false)} placement="end">
				<Offcanvas.Header closeButton>
					<Offcanvas.Title>Menu</Offcanvas.Title>
				</Offcanvas.Header>
				<Offcanvas.Body>
					<div className="d-grid gap-3">
						<Button 
							variant="link" 
							className="text-start p-0 text-dark text-decoration-none"
							onClick={() => scrollToSection('overview')}
						>
							Overview
						</Button>
						<Button 
							variant="link" 
							className="text-start p-0 text-dark text-decoration-none"
							onClick={() => scrollToSection('features')}
						>
							Features
						</Button>
						<Button 
							variant="link" 
							className="text-start p-0 text-dark text-decoration-none"
							onClick={() => scrollToSection('trusted')}
						>
							Made for
						</Button>
						<hr />
						{utilLinks.map(link => (
							<a key={link.href} href={link.href} className="btn btn-outline-secondary text-start">
								<i className={`${link.icon} me-2`}></i>
								{link.label}
							</a>
						))}
					</div>
				</Offcanvas.Body>
			</Offcanvas>
		</div>
	);
};

export default Home;
