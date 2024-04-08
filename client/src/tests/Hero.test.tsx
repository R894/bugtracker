import { expect, test } from 'vitest'
import { render, screen } from '@testing-library/react'

import Hero from '@/components/Hero/Hero'
import UserProvider from '@/context/UserContext'
import ProjectProvider from '@/context/ProjectContext'

test('renders BugTracker text', () => {
  render(
    <UserProvider>
      <ProjectProvider>
        <Hero />
      </ProjectProvider>
    </UserProvider>,
  )
  const bugTrackerText = screen.getByText('BugTracker')
  expect(bugTrackerText).toBeInTheDocument()
})

test('renders main message', () => {
  render(
    <UserProvider>
      <ProjectProvider>
        <Hero />
      </ProjectProvider>
    </UserProvider>,
  )
  const mainMessage = screen.getByText(
    'BugTracker centralizes communication, allowing teams to collaborate effortlessly and keep everyone on the same page.',
  )
  expect(mainMessage).toBeInTheDocument()
})
