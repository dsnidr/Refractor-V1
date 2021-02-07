import { css } from 'styled-components';

const breakpoints = {
	tiny: '480px',
	small: '768px',
	medium: '992px',
	large: '1200px',
	extralarge: '1500px',
};

export default Object.keys(breakpoints).reduce((accumulator, label) => {
	accumulator[label] = (...args) => css`
		@media (min-width: ${breakpoints[label]}) {
			${css(...args)};
		}
	`;
	return accumulator;
}, {});
