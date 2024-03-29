import { test, expect } from '@playwright/test'

test('has title', async ({ page }) => {
  await page.goto('/')
  await expect(page).toHaveTitle(/BugTracker/)
})

test('header buttons are visible', async ({ page }) => {
  const homeButton = page.getByRole('button', { name: 'Home' })
  const dashboardButton = page.getByRole('button', { name: 'Dashboard' })
  const loginButton = page.getByRole('button', { name: 'Login' })

  await page.goto('/')
  await expect(homeButton).toBeVisible()
  await expect(dashboardButton).toBeVisible()
  await expect(loginButton).toBeVisible()
})

test('header buttons redirect to correct pages', async ({ page }) => {
  const homeButton = page.getByRole('button', { name: 'Home' })
  const dashboardButton = page.getByRole('button', { name: 'Dashboard' })
  const loginButton = page.getByRole('button', { name: 'Login' })

  await page.goto('/')
  await homeButton.click()
  await expect(page).toHaveURL('/')

  await dashboardButton.click()
  await expect(page).toHaveURL('/login') // When user isn't logged in dashboard should redirect to login
  await page.goto('/')

  await loginButton.click()
  await expect(page).toHaveURL('/login')
})
