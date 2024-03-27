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
import Link from 'next/link'

const BugRow = ({ bug }: { bug: Bug }) => {
  return (
    <TableRow
      key={bug.id}
      sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
    >
      <TableCell component="th" scope="row">
        <b>
          <Link href={`/issue/${bug.id}`}>{bug.title}</Link>
        </b>
      </TableCell>
      <TableCell align="right">{bug.priority}</TableCell>
      <TableCell align="right">{bug.status}</TableCell>
      <TableCell align="right">{bug.asignee}</TableCell>
      <TableCell align="right">
        {moment(bug.createdAt).format('MMM Do YY')}
      </TableCell>
    </TableRow>
  )
}

const BugTable = ({ bugs }: { bugs: Bug[] }) => {
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
          {bugs &&
            bugs.length >= 1 &&
            bugs.map((row) => <BugRow key={row.id} bug={row} />)}
        </TableBody>
      </Table>
    </TableContainer>
  )
}

export default BugTable
