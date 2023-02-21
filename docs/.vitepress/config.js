export default {
    title: 'Luadoc',
    description: 'Lua documentation',
    themeConfig: {
        outline: [2,3],
        sidebar: [
          {
            text: 'Introduction',
            items: [
              { text: 'About Lua', link: '/01_intro/index' }
            ]
          },
          {
            text: 'Basic concepts',
            items: [
              { text: 'Introduction', link: '/02_basic_concepts/intro' },
              { text: 'Values and Types', link: '/02_basic_concepts/ch01' },
              { text: 'Error Handling', link: '/02_basic_concepts/ch03' },
              { text: 'Metatables and Metamethods', link: '/02_basic_concepts/ch04' },
              { text: 'Garbage Collection', link: '/02_basic_concepts/ch05' },
              { text: 'Coroutines', link: '/02_basic_concepts/ch06' },
            ]
          },
          {
            text: 'The Language',
            items: [
              { text: 'Introduction', link: '/03_the_language/intro' },
              { text: 'Lexical Conventions', link: '/03_the_language/ch01' },
              { text: 'Variables', link: '/03_the_language/ch02' },
              { text: 'Statements', link: '/03_the_language/ch03' },
              { text: 'Expressions', link: '/03_the_language/ch04' },
              { text: 'Visibility Rules', link: '/03_the_language/ch05' }
            ]
          },
          {
            text: 'The Application Program Interface',
            items: [
              { text: 'Introduction', link: '/04_API/intro' },
              { text: 'The Stack', link: '/04_API/ch01' },
              { text: 'C Closures', link: '/04_API/ch02' },
              { text: 'Registry', link: '/04_API/ch03' },
              { text: 'Error Handling in C', link: '/04_API/ch04' },
              { text: 'Handling Yields in C', link: '/04_API/ch05' },
              { text: 'Functions and Types', link: '/04_API/ch06' },
              { text: 'The Debug Interface', link: '/04_API/ch07' },
            ]
          },
          {
            text: 'The Auxiliary Library',
            items: [
              { text: 'Introduction', link: '/05_aux_lib/intro' },
              { text: 'Functions and Types', link: '/05_aux_lib/ch01' },
            ]
          },
          {
            text: 'The Standard Libraries',
            items: [
              { text: 'Introduction', link: '/06_standard_lib/intro' },
              { text: 'Basic Functions', link: '/06_standard_lib/ch01' },
              { text: 'Coroutine Manipulation', link: '/06_standard_lib/ch02' },
              { text: 'Modules', link: '/06_standard_lib/ch03' },
              { text: 'String Manipulation', link: '/06_standard_lib/ch04' },
              { text: 'UTF-8 Support', link: '/06_standard_lib/ch05' },
              { text: 'Table Manipulation', link: '/06_standard_lib/ch06' },
              { text: 'Mathematical Functions', link: '/06_standard_lib/ch07' },
              { text: 'Input and Output Facilities', link: '/06_standard_lib/ch08' },
              { text: 'Operating System Facilities', link: '/06_standard_lib/ch09' },
              { text: 'The Debug Library', link: '/06_standard_lib/ch10' },
            ]
          },
          {
            text: 'Lua Standalone',
            items: [
              { text: 'Lua Standalone', link: '/07_standalone/ch01' },
            ]
          },
          {
            text: 'Incompatibilities with the Previous Version',
            items: [
              { text: 'Introduction', link: '/08_incompatibilities/intro' },
              { text: 'Incompatibilities in the Language', link: '/08_incompatibilities/ch01' },
              { text: 'Incompatibilities in the Libraries', link: '/08_incompatibilities/ch02' },
              { text: 'Incompatibilities in the API', link: '/08_incompatibilities/ch03' },
            ]
          },
          {
            text: 'The Complete Syntax of Lua',
            items: [
              { text: 'The Complete Syntax of Lua', link: '/09_complete_syntax/ch01' },
            ]
          },
          {
            text: 'Full manual',
            items: [
              { text: 'Full manual in a single page', link: '/manual' },
            ]
          },
        ]
      }
}
