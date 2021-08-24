import React from 'react';
import ReactDOM from 'react-dom';
import CommentBox from './CommentBox';
import './index.css';

let data = [
];

ReactDOM.render(
	<CommentBox data={data}/>, 
	document.getElementById('root')
);
