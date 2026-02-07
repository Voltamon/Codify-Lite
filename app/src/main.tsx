import { render } from 'preact';
import { App } from './App';
import "@picocss/pico";
import './style.css';

render(<App />, document.getElementById('app')!);
