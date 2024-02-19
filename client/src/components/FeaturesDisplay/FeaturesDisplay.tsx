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
        padding: '16px'
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
          <Typography variant='h5'>Feature 1</Typography>
          <Typography>Write description here</Typography>
          <Typography>Write description here</Typography>
          <Typography>Write description here</Typography>
        </Feature>
        <Feature>
          <Apartment fontSize="large" />
          <Typography variant='h5'>Feature 2</Typography>
          <Typography>Write description here</Typography>
          <Typography>Write description here</Typography>
          <Typography>Write description here</Typography>
        </Feature>
        <Feature>
          <Apartment fontSize="large" />
          <Typography variant='h5'>Feature 3</Typography>
          <Typography>Write description here</Typography>
          <Typography>Write description here</Typography>
          <Typography>Write description here</Typography>
        </Feature>
      </Stack>
    </>
  )
}

export default FeaturesDisplay
