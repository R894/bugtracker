import { Lock, NoteAdd, Timeline } from '@mui/icons-material'
import { Card, Stack, Typography } from '@mui/material'
import { type ReactNode } from 'react'

const Feature = ({ children }: { children: ReactNode }) => {
  return (
    <Card
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        padding: '16px',
        maxWidth: '240px',
        minHeight: '240px',
      }}
    >
      {children}
    </Card>
  )
}

const FeaturesDisplay = () => {
  return (
    <Stack
      direction={'row'}
      gap={3}
      sx={{ flexWrap: 'wrap', justifyContent: 'center' }}
    >
      <Feature>
        <NoteAdd fontSize="large"></NoteAdd>
        <Typography variant="h5">Issue Creation</Typography>
        <Typography textAlign={'center'}>
          Easily create new issues with detailed information.
        </Typography>
      </Feature>
      <Feature>
        <Timeline fontSize="large"></Timeline>
        <Typography variant="h5">Issue Tracking</Typography>
        <Typography textAlign={'center'}>
          Keep track of the status, priority, and assignee for each issue.
        </Typography>
      </Feature>
      <Feature>
        <Lock fontSize="large"></Lock>
        <Typography variant="h5">Safety</Typography>
        <Typography textAlign={'center'}>
          Securely manage access with user authentication.
        </Typography>
      </Feature>
    </Stack>
  )
}

export default FeaturesDisplay
