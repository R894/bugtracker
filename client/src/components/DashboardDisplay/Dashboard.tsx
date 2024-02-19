import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material'

function createData(
  title: string,
  priority: string,
  status: string,
  assignee: string,
  createdAt: string,
) {
  return { title, priority, status, assignee, createdAt }
}

const rows = [
  createData('bug', 'critical', 'open', "user", 'may 15th'),
  createData('bug', 'critical', 'open', "user", 'may 15th'),
  createData('bug', 'critical', 'open', "user", 'may 15th'),
  createData('bug', 'critical', 'open', "user", 'may 15th'),
]

export default function BasicTable() {
  return (
    <div style={{padding:'16px'}}>
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
            {rows.map((row) => (
              <TableRow
                key={row.title}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.title}
                </TableCell>
                <TableCell align="right">{row.priority}</TableCell>
                <TableCell align="right">{row.status}</TableCell>
                <TableCell align="right">{row.assignee}</TableCell>
                <TableCell align="right">{row.createdAt}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  )
}
