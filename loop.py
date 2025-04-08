import whisper
from whisper.utils import get_writer

import glob
import os

model = whisper.load_model("tiny")
output_directory = "./"

for audio in glob.glob("*.mp4"):
    print(f"Transcribing: {audio}")
    result = model.transcribe(audio, fp16=False, language='en')
    sub_writer = get_writer("vtt", output_directory)
    sub_writer(result, audio)
    print(f"Done: {audio}")

print("All files transcribed!")

