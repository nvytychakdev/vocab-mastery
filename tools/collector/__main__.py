import typer
import questionary
import feed_collector
import words_translator
import words_collector

app = typer.Typer()

OPTIONS = {
    "Collect words definitions based on CERF dictionaries": words_collector.run,
    "Feed collected words to VocabMastery Database": feed_collector.run,
    "Translate all existing Words records from Database": words_translator.run,
}


@app.command()
def main():
    choice = questionary.select(
        "Select action to perform", choices=list(OPTIONS.keys())
    ).ask()

    if choice:
        OPTIONS[choice]()
    else:
        typer.echo("No option was selected. Closing the program. ")
    return


if __name__ == "__main__":
    app()
