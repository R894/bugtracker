import { expect, test } from 'vitest'
import { render, screen } from '@testing-library/react'
import FeaturesDisplay from '@/components/FeaturesDisplay/FeaturesDisplay'

test('renders FeaturesDisplay component correctly', () => {
  render(<FeaturesDisplay />)

  expect(screen.getByText('Issue Creation')).toBeInTheDocument()
  expect(screen.getByText('Issue Tracking')).toBeInTheDocument()
  expect(screen.getByText('Safety')).toBeInTheDocument()
})
