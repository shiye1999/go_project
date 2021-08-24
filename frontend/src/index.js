import React from 'react';
import ReactDOM from 'react-dom';
import CommentBox from './CommentBox';
import './index.css';

let data = [
  {author:'a',content:'这是第一条评论',date:'2021/8/24'},
  {author:'b',content:'这是第二条评论',date:'2021/8/24'}
];

ReactDOM.render(
	<CommentBox data={data}/>, 
	document.getElementById('root')
);
