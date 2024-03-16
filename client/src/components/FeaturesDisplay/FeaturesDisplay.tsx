import { Apartment } from '@mui/icons-material'
import { Card, Stack, Typography } from '@mui/material'
import { type ReactNode } from 'react'

const Feature = ({ children }: { children: ReactNode }) => {
  return (
    <Card
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        padding: '16px',
        maxWidth:'240px'
      }}
    >
      {children}
    </Card>
  )
}

const FeaturesDisplay = () => {
  return (
    <>
      <Stack direction={'row'} gap={3}>
        <Feature>
          <Apartment fontSize="large" />
          <Typography variant='h5'>Issue Creation</Typography>
          <Typography textAlign={'center'}>Easily create new issues with detailed information.</Typography>
        </Feature>
        <Feature>
          <Apartment fontSize="large" />
          <Typography variant='h5'>Issue Tracking</Typography>
          <Typography textAlign={'center'}>Keep track of the status, priority, and assignee for each issue.</Typography>
        </Feature>
        <Feature>
          <Apartment fontSize="large" />
          <Typography variant='h5'>Safe</Typography>
          <Typography textAlign={'center'}>Securely manage access with user authentication.</Typography>
        </Feature>
      </Stack>
    </>
  )
}

export default FeaturesDisplay
