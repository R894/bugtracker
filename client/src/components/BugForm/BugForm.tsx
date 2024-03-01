import { getBugById } from '@/api/bugService'
import { Bug } from '@/api/models/bug'
import { useUserContext } from '@/context/UserContext'
import { Subject } from '@mui/icons-material'
import { Box, Card, Divider, Stack, Typography } from '@mui/material'
import moment from 'moment'
import { useEffect, useState } from 'react'

const BugDisplay = ({ bugId }: { bugId: string }) => {
  const [bug, setBug] = useState<Bug | undefined>(undefined)
  const { token } = useUserContext()
  useEffect(() => {
    const getBug = async () => {
      const response = await getBugById(bugId, token)
      setBug(response)
    }
    getBug()
  }, [bugId, token])

  return (
    <>
      <Card
        variant="outlined"
        sx={{
          minWidth: '640px',
          minHeight: '72px',
          marginTop: '12px',
        }}
      >
        <Stack padding={1}>
          <Stack
            direction={'row'}
            gap={1}
            sx={{ display: 'flex', alignItems: 'center' }}
          >
            <Subject />
            <Typography variant="h5">
              {bug && bug.title.toUpperCase()}
            </Typography>
          </Stack>
          <Box paddingTop={1} sx={{ display: 'flex', justifyContent: 'space-between' }}>
            <Typography variant="subtitle2">
              Opened at {bug && moment(bug.createdAt).format('MMM Do YY')}
            </Typography>
            <Typography variant="subtitle2">
              Status: {bug && bug.status}
            </Typography>
            <Typography variant="subtitle2">
              Assigned to: {bug && bug.asignee}
            </Typography>
          </Box>
        </Stack>
        <Divider />
        <Box padding={1}>
          <Typography variant="body1">{bug && bug.description}</Typography>
          <Typography>{bug && bug.asignee}</Typography>
        </Box>
      </Card>
    </>
  )
}

export default BugDisplay
