import styles from "./index.less";
import {Redirect} from 'umi';
import cookie from 'react-cookies';

export default function IndexPage() {
  function authorized(): string | undefined {
    console.log(cookie.loadAll())
    return cookie.load("token")
  }

  console.log(authorized())

  return (
    <div>
      {authorized() ? <h1 className={styles.title}>Page index</h1> : <Redirect to="/login"/>}
    </div>
  );
}
