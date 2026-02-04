<?php

namespace App\Filament\Resources\Projects\Schemas;

use Filament\Forms\Components\DatePicker;
use Filament\Forms\Components\FileUpload;
use Filament\Forms\Components\TextInput;
use Filament\Forms\Components\RichEditor;
use Filament\Forms\Components\Toggle;
use Filament\Schemas\Schema;

class ProjectForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextInput::make('title')
                    ->required(),
                TextInput::make('slug')
                    ->required(),
                TextInput::make('short_description'),
                RichEditor::make('content')
                    ->label('Content')
                    ->columnSpanFull(),
                TextInput::make('repository_url')
                    ->url(),
                TextInput::make('live_url')
                    ->url(),
                FileUpload::make('thumbnail_url')
                    ->image()
                    ->directory('projects')
                    ->visibility('public'),
                DatePicker::make('started_at'),
                DatePicker::make('finished_at'),
                Toggle::make('is_featured')
                    ->required(),
            ]);
    }
}
