import React, { FC } from 'react';
import Layout from '../components/HomeLayout';

const sections = [
	{ id: 'overview', label: 'Overview' },
	{ id: 'stats', label: 'Statistics' },
	{ id: 'details', label: 'Details' },
];
  
const Home: FC = () => {
	const utilLinks = [
	  { href: '/account', label: 'Account' },
	  { href: '/settings', label: 'Settings' },
	  { href: '/help', label: 'Help' },
	];
  
	return (
	  <Layout leftLinks={utilLinks} menuItems={sections}>
		<section id="overview">
		  <h1>Overview</h1>
		  <p>General info about your account.</p>
		</section>
  
		<section id="stats">
		  <h1>Statistics</h1>
		  <p>Here are some interesting stats.</p>
		</section>
  
		<section id="details">
		  <h1>Details</h1>
		  <p>Detailed information and data table below.</p>
		</section>
	  </Layout>
	);
};

export default Home;
