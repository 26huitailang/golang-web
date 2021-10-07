import React, {ReactElement} from 'react';
import axios from 'axios';

export function Login (): ReactElement {
  axios.post('/login', {username: 'test', password: '123123'}).then((response: any) => {
    console.log(response.data)
  })
    .catch((error:any) => {
      console.log(error)
    })
  return (
    <div>Login</div>
  )
}
