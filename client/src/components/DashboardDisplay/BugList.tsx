import { Bug } from '@/api/models/bug'
import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material'
import moment from 'moment'

const BugList = ({bugs}:{bugs: Bug[]}) => {
  return (
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
            {bugs && bugs.length >= 1 && bugs.map((row) => (
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
  )
}

export default BugList
