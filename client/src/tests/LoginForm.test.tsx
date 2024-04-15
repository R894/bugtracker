import { expect, test } from 'vitest'
import { render, screen } from '@testing-library/react'
import LoginForm from '@/components/LoginForm/LoginForm'
import UserProvider from '@/context/UserContext'

test('renders login form correctly', () => {
  render(
    <UserProvider>
      <LoginForm />
    </UserProvider>,
  )

  expect(screen.getByText('Email')).toBeInTheDocument()
  expect(screen.getByText('Password')).toBeInTheDocument()
  expect(screen.queryByRole('button', { name: 'Login' })).toBeInTheDocument()
  expect(screen.getByText('Register here')).toBeInTheDocument()
})
