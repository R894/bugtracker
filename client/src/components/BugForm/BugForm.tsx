import { getBugById } from '@/api/bugService'
import { Bug } from '@/api/models/bug'
import { useUserContext } from '@/context/UserContext'
import { Box, Card, Stack, Typography } from '@mui/material'
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
      <Box>
        <Stack>
          <Typography variant="h6">{bug && bug.title}</Typography>
          <Typography>
            Opened at {bug && moment(bug.createdAt).format('MMM Do YY')}
          </Typography>
        </Stack>
        <Card
          variant="outlined"
          sx={{
            minWidth: '640px',
            minHeight: '320px',
            padding: '12px',
            marginTop: '12px',
          }}
        >
          <Typography>{bug && bug.description}</Typography>
          <Typography>{bug && bug.asignee}</Typography>
        </Card>
      </Box>
    </>
  )
}

export default BugDisplay
