import { useState } from 'react';

function useValue(initValue) {
  const [value, setValue] = useState(initValue);
  const handleValue = e => {
    setValue(e.target.value);
  };
  return {
    value,
    handleValue
  };
}

export default useValue;
