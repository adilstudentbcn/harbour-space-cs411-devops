# Reflection — introduction-to-builds

## Stretch Task: Strip the Binary
I built the stripped binary and compared the sizes.
 The `main-stripped` file is much smaller because the compiler removed all the extra human-readable debugging text. 



## What did I do?
I went into my app folder and used a command (`go build`) to turn my raw code into a real, working program. Then, I turned the program on. 

## What was most surprising?
I was surprised by how much I could shrink the app's file size just by using a special command to delete the hidden "debugging" text. It was interesting to see how much extra space is taken up by text that is only there to help humans read errors.

## What's still unclear?
I understand that CGO_ENABLED=0 fixes the crash by packing all the code into one file. But I still don't understand how Go decides what parts of the operating system it needs to borrow from in the first place.