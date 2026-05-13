import React from 'react';
import { Select } from 'antd';

const normalizeOptions = (options) =>
  (Array.isArray(options) ? options : []).map((option) => ({
    label: option?.label ?? option?.text ?? option?.name ?? option?.value,
    value: option?.value,
    disabled: option?.disabled === true,
  }));

function AppSelect({
  className = '',
  options = [],
  value,
  placeholder,
  disabled,
  name,
  search,
  clearable,
  multiple,
  fluid = false,
  compact = false,
  noResultsMessage,
  onChange,
  ...props
}) {
  const nextClassName = ['router-ui-select', fluid ? 'fluid' : '', className]
    .filter(Boolean)
    .join(' ');
  const mode = multiple === true ? 'multiple' : undefined;
  const normalizedOptions = normalizeOptions(options);
  const size = compact ? 'small' : undefined;

  const filterOption =
    typeof search === 'function'
      ? (inputValue, option) =>
          search(
            normalizedOptions.map((item) => ({
              key: item.value,
              value: item.value,
              text: item.label,
            })),
            inputValue,
          ).some((item) => item?.value === option?.value)
      : undefined;

  const handleChange = (nextValue, option) => {
    if (typeof onChange === 'function') {
      onChange(null, {
        name,
        value: nextValue,
        option,
      });
    }
  };

  return (
    <Select
      {...props}
      className={nextClassName}
      options={normalizedOptions}
      value={value}
      placeholder={placeholder}
      disabled={disabled}
      showSearch={search === true || typeof search === 'function'}
      filterOption={filterOption}
      allowClear={clearable === true}
      mode={mode}
      size={size}
      notFoundContent={noResultsMessage}
      onChange={handleChange}
    />
  );
}

export default AppSelect;
