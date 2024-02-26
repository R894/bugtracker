import { getBugs } from '@/api/bugService'
import { Bug } from '@/api/models/bug'
import { UserContext, UserContextType } from '@/context/UserContext'
import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material'
import { useContext, useEffect, useState } from 'react'
import moment from 'moment'

const Dashboard = () => {
  const { token, updateUserProjects } = useContext(
    UserContext,
  ) as UserContextType
  const [bugs, setBugs] = useState<Bug[]>([])

  useEffect(() => {
    updateUserProjects()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    async function fetchBugs() {
      if (!token) return
      const response: Bug[] = await getBugs(token)
      setBugs(response)
      console.log(response)
    }
    fetchBugs()
  }, [token])

  return (
    <div style={{ padding: '16px' }}>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 1080 }} aria-label="Bug Table">
          <TableHead>
            <TableRow>
              <TableCell>Title</TableCell>
              <TableCell align="right">Priority</TableCell>
              <TableCell align="right">Status</TableCell>
              <TableCell align="right">Assignee</TableCell>
              <TableCell align="right">Date Created</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {bugs.map((row) => (
              <TableRow
                key={row.id}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.title}
                </TableCell>
                <TableCell align="right">{row.priority}</TableCell>
                <TableCell align="right">{row.status}</TableCell>
                <TableCell align="right">{row.asignee}</TableCell>
                <TableCell align="right">
                  {moment(row.createdAt).format('MMM Do YY')}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  )
}

export default Dashboard
