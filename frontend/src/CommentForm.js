import React, { Component } from 'react';

class CommentForm extends Component {
  handleSubmit(e) {
    e.preventDefault();
    let content = this.refs.content.value,
        author = this.refs.author.value,
        date = new Date().toLocaleString(),
        warning = this.refs.warning;
    
    if(!author) {
      warning.innerHTML = '* 姓名不能为空';
      return null;
    }else if(!content) {
      warning.innerHTML = '* 评论不能为空';
      return null;
    }else {
      warning.innerHTML = '';
    }
    
    this.props.handleCommentSubmit({content, author, date});
  }

  render() {
    return (
       <form className='submitform' onSubmit={(e) => this.handleSubmit(e)}>
        <span className='warning' ref='warning'></span>
        <div className='form-row'>
          <input type="text" placeholder='姓名' ref='author'/>
        </div>
        <div className='form-row'>
          <textarea placeholder='评论' ref='content'></textarea>
        </div>
        <div className='form-row'>
          <button>发表</button>
        </div>
      </form>
    );
  }
}

export default CommentForm;
