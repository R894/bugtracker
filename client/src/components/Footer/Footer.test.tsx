import { expect, test } from 'vitest'
import { render, screen } from '@testing-library/react'
import Footer from './Footer';

test('renders footer text correctly', () => {
    render(<Footer />);
    const footerText = screen.getByText('Copyright © BugTracker');
    expect(footerText).toBeInTheDocument();
});