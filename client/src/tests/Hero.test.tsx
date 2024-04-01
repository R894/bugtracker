import { expect, test } from 'vitest'
import { render, screen } from '@testing-library/react'

import Hero from '@/components/Hero/Hero'

test('renders BugTracker text', () => {
  render(<Hero />)
  const bugTrackerText = screen.getByText('BugTracker')
  expect(bugTrackerText).toBeInTheDocument()
})

test('renders main message', () => {
  render(<Hero />)
  const mainMessage = screen.getByText(
    'BugTracker centralizes communication, allowing teams to collaborate effortlessly and keep everyone on the same page.',
  )
  expect(mainMessage).toBeInTheDocument()
})
