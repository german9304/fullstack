function Comment({ data }) {
  return (
    <div className='comment comment-body'>
      <p className='body'>{data.body}</p>
    </div>
  );
}

export default Comment;
