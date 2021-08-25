import styles from './index.less';
import Login from './Login'

export default function IndexPage() {
  let authorized = false
  return (
    <div>
      {authorized ? <h1 className={styles.title}>Page index</h1> : <Login/>}
    </div>
  );
}
