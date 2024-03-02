import { TextField } from '@mui/material'
import { useState } from 'react'

const CommentForm = () => {
  const [commentText, setCommentText] = useState('')
  return (
    <TextField
      sx={{ p: '20px 0' }}
      multiline
      fullWidth
      minRows={2}
      id="outlined-multilined"
      placeholder="Write comment here..."
      value={commentText}
      onChange={(e) => {
        setCommentText(e.target.value)
      }}
    />
  )
}

export default CommentForm
