import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Modal from '@mui/material/Modal'
import { TextField, Typography } from '@mui/material'
import { useContext, useState } from 'react'
import { UserContext, UserContextType } from '@/context/UserContext'
import { createBug } from '@/api/bugService'

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  boxShadow: 24,
  display: 'flex',
  flexDirection: 'column',
  p: 4,
}

const CreateBugModal = ({projectId}: {projectId: string}) => {
  const [open, setOpen] = useState(false)
  const [bugName, setBugName] = useState('')
  const [bugDescription, setBugDescription] = useState('')
  const {token} = useContext(UserContext) as UserContextType
  const handleOpen = () => setOpen(true)
  const handleClose = () => setOpen(false)

  const postProject = async () => {
    const response = await createBug({title: bugName, description: bugDescription, projectId: projectId}, token)
    console.log(response)
    handleClose()
  }

  return (
    <div>
      <Button onClick={handleOpen}>New Issue</Button>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography variant="h6" sx={{ textAlign: 'center' }}>
            New Issue
          </Typography>
          <TextField
            variant="standard"
            value={bugName}
            onChange={(e) => setBugName(e.target.value)}
            label="Name"
          />
          <TextField
            variant="standard"
            value={bugDescription}
            onChange={(e) => setBugDescription(e.target.value)}
            label="Description"
          />
          <Button onClick={postProject}>Create</Button>
        </Box>
      </Modal>
    </div>
  )
}

export default CreateBugModal
