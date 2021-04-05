export default ({ properties }) => (key = 'MethodName:asc') => {
  return properties(['MethodName', 'TTL'])(key);
};
